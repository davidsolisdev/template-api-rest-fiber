package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id                   uint   `json:"id"`
	Name                 string `json:"name" validate:"required,min=3"`
	LastName             string `json:"lastName" validate:"required,min=3"`
	Email                string `json:"email" validate:"required,min=5" gorm:"column:email"`
	Password             string `validate:"required,min=8"`
	LastPassword         string
	Role                 string
	ConfirmedEmail       bool
	ConfirmedEmailSecret string
	CodeRecoverPassword  string
	Created              time.Time
	Updated              time.Time
	Deleted              gorm.DeletedAt `gorm:"index"`
}
