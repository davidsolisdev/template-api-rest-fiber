package database

import (
	"os"

	"github.com/davidsolisdev/template-api-rest-fiber/helpers"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	//"github.com/davidalarcon68/ApiRest/models"
)

var DBMSSql *gorm.DB

func ConnectMsSql() {
	var err error
	var dsn string = "Server=" + os.Getenv("HOSTDB") + ";Database=" + os.Getenv("DB") + ";User Id=" + os.Getenv("USERDB") + ";Password=" + os.Getenv("PASSWORDDB") + ";Encrypt=disable;TrustServerCertificate=False"

	DBMSSql, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		helpers.ErrorDatabase("Sql Server", err)
	}

	//db.AutoMigrate(&models.User{})
}