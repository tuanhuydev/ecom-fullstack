package models

import "time"

type Collection struct {
	ID          string     `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()" validate:"required"`
	Name        string     `json:"name" gorm:"type:varchar(255);not null" validate:"required"`
	Description string     `json:"description" gorm:"type:text"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}
