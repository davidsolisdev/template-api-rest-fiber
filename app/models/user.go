package models

import "time"

type User struct {
	Id             uint   `json:"id"`
	Name           string `json:"name" validate:"required,min=3"`
	LastName       string `json:"lastName" validate:"required,min=3"`
	Email          string `json:"email" validate:"required,min=5" gorm:"column:email"`
	Password       string `validate:"required,min=8"`
	Role           string
	ConfirmedEmail bool
	Created        time.Time
}
