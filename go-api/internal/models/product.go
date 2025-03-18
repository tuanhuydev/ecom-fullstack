package models

import "time"

type Product struct {
	ID          string     `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string     `json:"name" gorm:"type:varchar(50);not null" validate:"required,min=3,max=50"`
	Price       float64    `json:"price" gorm:"type:decimal;not null" validate:"required,gt=0"`
	Description string     `json:"description" gorm:"type:varchar(255);not null"`
	Quantity    int        `json:"quantity" gorm:"type:int;not null" validate:"required,gte=0"`
	Thumbnail   *string    `json:"thumbnail" gorm:"type:varchar(255)" validate:"omitempty,url"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}
