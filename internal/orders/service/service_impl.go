package service

import (
	"fmt"
	"test-be-dbo/internal/helpers"
	"test-be-dbo/internal/models"
	orderRepository "test-be-dbo/internal/orders/repository"

	"github.com/google/uuid"
)

type orderService struct {
	orderRepo orderRepository.Repository
}

func NewOrderService(orderRepo orderRepository.Repository) Service {
	return &orderService{orderRepo: orderRepo}
}

func (o *orderService) ProductIDExists(request models.OrderRequest) []string {
	var invalidProductIDs []string
	for _, product := range request.Products {
		_, err := o.orderRepo.ProductIDExists(product.ID)
		if err != nil {
			invalidProductIDs = append(invalidProductIDs, fmt.Sprintf("product with ID %s not found", product.ID))
		}
	}
	return invalidProductIDs
}

func (o *orderService) CreateOrder(request models.OrderRequest) (models.OrderResponse, error) {
	var (
		orderResponse models.OrderResponse
		orderDetails  []models.OrderDetail
	)

	order := models.Order{
		CustomerID: helpers.ParseUUID(request.CustomerID),
		TotalItems: uint16(len(request.Products)),
		Status:     models.UNPAID, // Set status default unpaid,
		CreatedBy:  request.CreatedBy,
	}

	// Calculate the total price and populate order details
	for _, productReq := range request.Products {
		product, err := o.orderRepo.GetProductByID(productReq.ID)
		if err != nil {
			return orderResponse, err
		}

		if productReq.Qty > product.Qty {
			productReq.Qty = product.Qty
		}

		totalPrice := product.Price * float64(productReq.Qty)

		orderDetail := models.OrderDetail{
			ProductID:   helpers.ParseUUID(productReq.ID),
			ProductName: product.Name,
			UnitPrice:   product.Price,
			Qty:         productReq.Qty,
			TotalPrice:  totalPrice,
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	order.OrderDetails = orderDetails
	order.TotalPrice = calculateOrderTotalPrice(orderDetails)

	result, err := o.orderRepo.CreateOrder(order)
	if err != nil {
		return orderResponse, err
	}

	response := models.OrderResponse{
		ID:          result.ID.String(),
		CustomerID:  request.CustomerID,
		Status:      result.Status,
		TotalItems:  result.TotalItems,
		TotalPrice:  result.TotalPrice,
		OrderDetail: make([]models.OrderDetailResponse, len(result.OrderDetails)),
		CreatedAt:   &result.CreatedAt,
		CreatedBy:   result.CreatedBy,
	}

	for i, orderDetail := range result.OrderDetails {
		response.OrderDetail[i] = models.OrderDetailResponse{
			ID:          result.ID.String(),
			ProductID:   orderDetail.ProductID.String(),
			ProductName: orderDetail.ProductName,
			UnitPrice:   orderDetail.UnitPrice,
			Qty:         orderDetail.Qty,
			TotalPrice:  orderDetail.TotalPrice,
		}
	}

	return response, nil

}

func calculateOrderTotalPrice(orderDetails []models.OrderDetail) float64 {
	var totalPrice float64
	for _, od := range orderDetails {
		totalPrice += od.TotalPrice
	}
	return totalPrice
}

func (o *orderService) OrderIDExist(id uuid.UUID) (bool, error) {
	_, err := o.orderRepo.OrderIDExist(id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (o *orderService) OrderHasPaid(id uuid.UUID) (bool, error) {
	_, err := o.orderRepo.OrderHasPaid(id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (o *orderService) DeleteOrder(id uuid.UUID) error {
	if err := o.orderRepo.DeleteOrder(id); err != nil {
		return err
	}

	return nil
}

func (o *orderService) UpdateOrder(request models.OrderRequest, id uuid.UUID) (models.OrderResponse, error) {
	var (
		orderResponse models.OrderResponse
		orderDetails  []models.OrderDetail
	)

	order := models.Order{
		CustomerID: helpers.ParseUUID(request.CustomerID),
		TotalItems: uint16(len(request.Products)),
		Status:     models.UNPAID, // Set status default unpaid,
		CreatedBy:  request.CreatedBy,
	}

	// Calculate the total price and populate order details
	for _, productReq := range request.Products {
		product, err := o.orderRepo.GetProductByID(productReq.ID)
		if err != nil {
			return orderResponse, err
		}

		if productReq.Qty > product.Qty {
			productReq.Qty = product.Qty
		}

		totalPrice := product.Price * float64(productReq.Qty)

		orderDetail := models.OrderDetail{
			ProductID:   helpers.ParseUUID(productReq.ID),
			ProductName: product.Name,
			UnitPrice:   product.Price,
			Qty:         productReq.Qty,
			TotalPrice:  totalPrice,
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	order.OrderDetails = orderDetails
	order.TotalPrice = calculateOrderTotalPrice(orderDetails)

	result, err := o.orderRepo.UpdateOrder(order, id)
	if err != nil {
		return orderResponse, err
	}

	response := models.OrderResponse{
		ID:          result.ID.String(),
		CustomerID:  request.CustomerID,
		Status:      result.Status,
		TotalItems:  result.TotalItems,
		TotalPrice:  result.TotalPrice,
		OrderDetail: make([]models.OrderDetailResponse, len(result.OrderDetails)),
		UpdatedAt:   &result.UpdatedAt,
		UpdatedBy:   result.UpdatedBy,
	}

	for i, orderDetail := range result.OrderDetails {
		response.OrderDetail[i] = models.OrderDetailResponse{
			ID:          result.ID.String(),
			ProductID:   orderDetail.ProductID.String(),
			ProductName: orderDetail.ProductName,
			UnitPrice:   orderDetail.UnitPrice,
			Qty:         orderDetail.Qty,
			TotalPrice:  orderDetail.TotalPrice,
		}
	}

	return response, nil

}
