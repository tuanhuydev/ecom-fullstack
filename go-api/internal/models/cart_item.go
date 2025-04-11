package models

import "time"

type CartItem struct {
	ID        string     `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	AccountID string     `json:"accountId" gorm:"type:uuid;not null"`
	ProductID string     `json:"productId" gorm:"type:uuid;not null"`
	Quantity  int        `json:"quantity" gorm:"type:int;not null" validate:"required,gte=0"`
	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}
