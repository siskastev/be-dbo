package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type status string

const (
	PAID   status = "paid"
	UNPAID status = "unpaid"
)

func (ct *status) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*ct = status(v)
	case string:
		*ct = status(v)
	default:
		return fmt.Errorf("unsupported type for status: %T", value)
	}
	return nil
}

func (ct status) Value() (driver.Value, error) {
	return string(ct), nil
}

type Order struct {
	gorm.Model
	ID           uuid.UUID      `gorm:"type:uuid;column:id;primaryKey;default:gen_random_uuid()"`
	CustomerID   uuid.UUID      `gorm:"column:customer_id;not null"`
	TotalItems   uint16         `gorm:"column:total_items;not null"`
	TotalPrice   float64        `gorm:"column:total_price;not null;"`
	Status       status         `gorm:"column:status;not null;"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy    string         `gorm:"column:created_by;type:varchar(100);not null"`
	UpdatedBy    string         `gorm:"column:updated_by;type:varchar(100);not null"`
	OrderDetails []OrderDetail  `gorm:"foreignKey:OrderID"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;"`
	Customer     Customer       `gorm:"foreignKey:CustomerID"`
}

type OrderRequest struct {
	CustomerID string                `json:"customer_id" binding:"required"`
	Products   []ProductOrderRequest `json:"products" binding:"required,dive"`
	CreatedBy  string                `json:"created_by,omitempty"`
	UpdatedBy  string                `json:"updated_by,omitempty"`
}

type ProductOrderRequest struct {
	ID  string `json:"product_id" binding:"required"`
	Qty uint16 `json:"qty" binding:"required,numeric,min=1,max=100"`
}

type OrderResponse struct {
	ID          string                `json:"order_id"`
	CustomerID  string                `json:"customer_id"`
	Status      status                `json:"status"`
	TotalItems  uint16                `json:"total_items"`
	TotalPrice  float64               `json:"total_price"`
	OrderDetail []OrderDetailResponse `json:"products"`
	CreatedAt   *time.Time            `json:"created_at,omitempty"`
	CreatedBy   string                `json:"created_by,omitempty"`
	UpdatedAt   *time.Time            `json:"updated_at,omitempty"`
	UpdatedBy   string                `json:"updated_by,omitempty"`
}

type OrderDetailResponse struct {
	ID          uint    `json:"order_id"`
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	UnitPrice   float64 `json:"unit_price"`
	Qty         uint16  `json:"qty"`
	TotalPrice  float64 `json:"total_price"`
}

type ManageOrderResponse struct {
	ID           string                `json:"order_id"`
	CustomerID   string                `json:"customer_id"`
	CustomerName string                `json:"customer_name"`
	Status       status                `json:"status"`
	TotalItems   uint16                `json:"total_items"`
	TotalPrice   float64               `json:"total_price"`
	CreatedAt    time.Time             `json:"created_at,omitempty"`
	CreatedBy    string                `json:"created_by,omitempty"`
	UpdatedAt    time.Time             `json:"updated_at,omitempty"`
	UpdatedBy    string                `json:"updated_by,omitempty"`
	OrderDetail  []OrderDetailResponse `json:"orders_detail,omitempty"`
}

type FilterOrders struct {
	ID         string `form:"id"`
	Name       string `form:"customer_name"`
	TotalItems int    `form:"total_items"`
}
