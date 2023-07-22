package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gender string

const (
	FEMALE gender = "female"
	MALE   gender = "male"
)

type Customer struct {
	gorm.Model
	ID        uuid.UUID      `gorm:"type:uuid;column:id;primaryKey;default:gen_random_uuid()"`
	Address   string         `gorm:"column:address;not null"`
	Name      string         `gorm:"column:name;type:varchar(100);not null"`
	Email     string         `gorm:"column:email;type:varchar(100);unique;not null"`
	Phone     string         `gorm:"column:phone;type:varchar(100);not null"`
	Gender    gender         `gorm:"column:gender;not null"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	CreatedBy string         `gorm:"column:created_by;type:varchar(100);not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	UpdatedBy string         `gorm:"column:updated_by;type:varchar(100);not null"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;"`
}

type CustomerRequest struct {
	Name      string `json:"name" binding:"required,min=2,max=100"`
	Email     string `json:"email" binding:"required,email,max=100"`
	Address   string `json:"address" binding:"required"`
	Gender    gender `json:"gender" binding:"required,eq=female|eq=male"`
	Phone     string `json:"phone" binding:"required,numeric,min=10,max=14"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type CustomerResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Address   string     `json:"address"`
	Gender    gender     `json:"gender"`
	Phone     string     `json:"phone"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	CreatedBy string     `json:"created_by,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy string     `json:"updated_by,omitempty"`
}
