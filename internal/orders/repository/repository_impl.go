package repository

import (
	"test-be-dbo/internal/helpers"
	"test-be-dbo/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) Repository {
	return &orderRepository{db: db}
}

func (o *orderRepository) ProductIDExists(id string) (bool, error) {
	var product models.Product
	if err := o.db.Where("id = ? AND qty > 0", id).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return product.ID.String() != "", nil

}

func (o *orderRepository) CreateOrder(order models.Order) (models.Order, error) {
	var createdOrder models.Order
	err := o.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		for _, orderDetail := range order.OrderDetails {
			product, err := o.getProductByID(tx, orderDetail.ProductID)
			if err != nil {
				return err
			}

			product.Qty -= orderDetail.Qty
			if err := tx.Where("id = ?", product.ID).Updates(&product).Error; err != nil {
				return err
			}
		}

		createdOrder = order
		return nil
	})

	if err != nil {
		return createdOrder, err
	}

	return createdOrder, nil
}

func (o *orderRepository) GetProductByID(id string) (models.Product, error) {
	var product models.Product
	if err := o.db.Where("id = ?", id).First(&product).Error; err != nil {
		return product, nil
	}

	return product, nil
}

func (o *orderRepository) getProductByID(tx *gorm.DB, productID uuid.UUID) (*models.Product, error) {
	var product models.Product
	if err := tx.Where("id = ?", productID).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (o *orderRepository) getOrderById(tx *gorm.DB, id uuid.UUID) (models.Order, error) {
	var order models.Order
	if err := tx.Preload("OrderDetails").First(&order, id).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (o *orderRepository) GetOrderByID(id uuid.UUID) (models.Order, error) {
	var order models.Order
	if err := o.db.Preload("OrderDetails").Unscoped().Preload("Customer").First(&order, id).Error; err != nil {
		return order, err
	}

	return order, nil
}

func (o *orderRepository) DeleteOrder(id uuid.UUID) error {
	if err := o.db.Delete(&models.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) OrderIDExist(id uuid.UUID) (bool, error) {
	var order models.Order
	if err := o.db.Where(models.Order{ID: id}).First(&order).Error; err != nil {
		return false, err
	}
	return order.ID.String() != "", nil
}

func (o *orderRepository) UpdateOrder(order models.Order, id uuid.UUID) (models.Order, error) {
	var updateOrder models.Order

	order.ID = id

	err := o.db.Transaction(func(tx *gorm.DB) error {

		// Fetch the existing order from the database
		existingOrder, err := o.getOrderById(tx, id)
		if err != nil {
			return err
		}

		// Reverse the stock changes for existing OrderDetail records
		for _, existingOrderDetail := range existingOrder.OrderDetails {
			product, err := o.getProductByID(tx, existingOrderDetail.ProductID)
			if err != nil {
				return err
			}
			product.Qty += existingOrderDetail.Qty
			if err := tx.Model(&product).Updates(&product).Error; err != nil {
				return err
			}
		}

		// Delete all existing products associated with the order
		if err := tx.Where(models.OrderDetail{OrderID: id}).Delete(&models.OrderDetail{}).Error; err != nil {
			return err
		}

		// Update the order details and other fields of the order itself
		if err := o.db.Model(models.Order{ID: id}).Updates(&order).Error; err != nil {
			return err
		}

		// Create new OrderDetail records with the correct OrderID
		for _, orderDetail := range order.OrderDetails {
			product, err := o.getProductByID(tx, orderDetail.ProductID)
			if err != nil {
				return err
			}
			product.Qty -= orderDetail.Qty
			if err := tx.Model(&product).Updates(&product).Error; err != nil {
				return err
			}

			// Create new OrderDetail record with the correct OrderID
			orderDetail.OrderID = id
			if err := tx.Create(&orderDetail).Error; err != nil {
				return err
			}
		}

		updateOrder = order
		return nil
	})

	if err != nil {
		return updateOrder, err
	}

	return updateOrder, nil
}

func (o *orderRepository) OrderHasPaid(id uuid.UUID) (bool, error) {
	var order models.Order
	if err := o.db.Where(models.Order{ID: id, Status: models.PAID}).First(&order).Error; err != nil {
		return false, err
	}
	return order.ID.String() != "", nil
}

func (o *orderRepository) GetAll(paginationParams helpers.PaginationParams, filters models.FilterOrders) ([]models.Order, int64, error) {
	var orders []models.Order

	query := o.db.Model(models.Order{}).Unscoped().Preload("Customer")

	if filters.Name != "" {
		query = query.Joins("JOIN customers ON orders.customer_id = customers.id").
			Where("customers.name LIKE ?", "%"+filters.Name+"%")
	}

	if filters.ID != "" {
		query = query.Where("id = ?", filters.ID)
	}

	if filters.Status != "" {
		query = query.Where("status = ?", filters.Status)
	}

	if filters.TotalItems > 0 {
		query = query.Where("total_items = ?", filters.TotalItems)
	}

	var totalRecords int64
	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	offset := paginationParams.GetOffset()
	limit := paginationParams.PageSize

	if err := query.Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, totalRecords, nil
}
