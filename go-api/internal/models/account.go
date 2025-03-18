package models

import "time"

type Account struct {
	ID        string     `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserId    string     `gorm:"type:uuid;not null"`
	Password  string     `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"autoCreateTime"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"index"`
}
