package utils

import (
	"log"
)

func ErrorDatabase(dbType string, err error) {
	log.Fatal("Failed to connect to database: "+dbType+" \n", err.Error())
}
