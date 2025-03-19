package models

import "time"

type Order struct {
	ID            string     `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	AccountID     string     `json:"accountId" gorm:"type:uuid;not null"`
	TotalPrice    float64    `json:"totalPrice" gorm:"type:decimal;not null"`
	Status        string     `json:"status" gorm:"type:varchar(50);not null" validate:"required,min=3,max=50"`
	PaymentMethod string     `json:"paymentMethod" gorm:"type:varchar(50);not null" validate:"required,min=3,max=50"`
	CreatedAt     time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt     *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}
