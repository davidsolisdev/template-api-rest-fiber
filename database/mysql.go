package database

import (
	"os"

	"github.com/davidsolisdev/template-api-rest-fiber/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	//"github.com/davidalarcon68/ApiRest/models"
)

var DBMySql *gorm.DB

func ConnectMySql() {
	var err error
	var dsn string = os.Getenv("USERDB") + ":" + os.Getenv("PASSWORDDB") + "@tcp(" + os.Getenv("HOSTDB") + ":" + os.Getenv("PORTDB") + ")/" + os.Getenv("DB") + "?charset=utf8mb4&parseTime=True&loc=Local"

	DBMySql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		helpers.ErrorDatabase("MySql", err)
	}

	//db.AutoMigrate(&models.User{})
}
