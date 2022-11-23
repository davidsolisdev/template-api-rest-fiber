package entity

import "time"

type User struct {
	Id                   uint   `json:"id"`
	Name                 string `json:"name" validate:"required,min=3"`
	LastName             string `json:"lastName" validate:"required,min=3"`
	Email                string `json:"email" validate:"required,min=5" gorm:"column:email"`
	LastEmail            string `json:"last_email"`
	Password             string `validate:"required,min=8"`
	LastPassword         string
	Role                 string
	ConfirmedEmail       bool
	ConfirmedEmailSecret string
	CodeRecoverPassword  string
	Created              time.Time
	Updated              time.Time
	Deleted              time.Time
}

/*
CREATE TABLE users(
   id 			           SERIAL 	   PRIMARY KEY  NOT NULL,
   name                    CHAR(50)    NOT NULL,
   last_name               CHAR(50)    NOT NULL,
   email                   CHAR(50)    NOT NULL,
   last_email              CHAR(200),
   password     		   CHAR(200)    NOT NULL,
   last_password           CHAR(200),
   role                    CHAR(50)    NOT NULL,
   confirmed_email         BOOL        NOT NULL,
   confirmed_email_secret  CHAR(75)    NOT NULL,
   code_recover_password   CHAR(75),
   created                 timestamp   NOT NULL,
   updated                 timestamp   NOT NULL,
   deleted                 timestamp
);
*/
