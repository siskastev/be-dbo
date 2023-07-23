package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderDetail struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement"`
	OrderID     uuid.UUID `gorm:"column:order_id;"`
	ProductID   uuid.UUID `gorm:"column:product_id;"`
	ProductName string    `gorm:"column:product_name;"`
	UnitPrice   float64   `gorm:"column:unit_price;not null"`
	Qty         uint16    `gorm:"column:qty;not null"`
	TotalPrice  float64   `gorm:"column:total_price;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (OrderDetail) TableName() string {
	return "orders_detail"
}
