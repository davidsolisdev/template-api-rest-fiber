package database

import (
	"os"

	"github.com/davidsolisdev/template-api-rest-fiber/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	//"github.com/davidalarcon68/ApiRest/models"
)

var DBMySql *gorm.DB

func ConnectMySql() {
	var err error
	var dsn string = os.Getenv("USER_DB") + ":" + os.Getenv("PASSWORD_DB") + "@tcp(" + os.Getenv("HOST_DB") + ":" + os.Getenv("PORT_DB") + ")/" + os.Getenv("DB") + "?charset=utf8mb4&parseTime=True&loc=Local"

	DBMySql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		utils.ErrorDatabase("MySql", err)
	}

	//DBMySql.AutoMigrate(&models.User{})
}
