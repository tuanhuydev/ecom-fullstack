package models

import "time"

type CollectionProduct struct {
	ProductID    string     `json:"productId" gorm:"type:uuid;primaryKey" validate:"required"`
	CollectionID string     `json:"collectionId" gorm:"type:uuid;primaryKey" validate:"required"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}
