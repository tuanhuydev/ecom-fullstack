package models

import "time"

type OrderProduct struct {
	OrderID   string     `json:"orderId" gorm:"type:uuid;primaryKey"`
	ProductID string     `json:"productId" gorm:"type:uuid;primaryKey"`
	Quantity  int        `json:"quantity" gorm:"type:int;not null;default:1" validate:"required,gte=0"`
	Price     float64    `json:"price" gorm:"type:decimal;not null" validate:"required,gt=0"`
	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}
