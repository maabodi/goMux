package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    int            `gorm:"primary_key; AUTO_INCREMENT" json:"-"`
	Email     string         `json:"email" form:"email"`
	Role      string         `json:"role" form:"role"`
	Password  string         `json:"password" form:"password"`
}
