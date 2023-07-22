package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;column:id;primaryKey;default:gen_random_uuid()"`
	Name      string    `gorm:"column:name;type:varchar(100);not null"`
	Qty       uint16    `gorm:"column:qty;not null"`
	Price     float64   `gorm:"column:price;not null;"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
