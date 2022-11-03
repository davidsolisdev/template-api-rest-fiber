package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" validate:"required,min=3"`
	LastName string `json:"lastName" validate:"required,min=3" gorm:"column:lastName"`
	Email    string `json:"email" validate:"required,min=5" gorm:"column:email"`
	Password string `json:"password" validate:"required,min=8" gorm:"column:password"`
	Role     string `json:"role" gorm:"column:role"`
	Created  time.Time
	Updated  time.Time
	Deleted  gorm.DeletedAt `gorm:"index"`
}
