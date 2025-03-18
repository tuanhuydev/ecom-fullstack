package models

import "time"

type User struct {
	ID        string     `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Account   Account    `gorm:"foreignKey:UserId"`
	Name      string     `json:"name" gorm:"type:varchar(255);not null"`
	Email     string     `json:"email" gorm:"type:varchar(50);not null"`
	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"autoCreateTime"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"index"`
}
