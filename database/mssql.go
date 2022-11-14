package database

import (
	"os"

	"github.com/davidsolisdev/template-api-rest-fiber/utils"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	//"github.com/davidalarcon68/ApiRest/models"
)

var DBMSSql *gorm.DB

func ConnectMsSql() {
	var err error
	var dsn string = "Server=" + os.Getenv("HOST_DB") + ";Database=" + os.Getenv("DB") + ";User Id=" + os.Getenv("USER_DB") + ";Password=" + os.Getenv("PASSWORD_DB") + ";Encrypt=disable;TrustServerCertificate=False"

	DBMSSql, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		utils.ErrorDatabase("Sql Server", err)
	}

	//DBMSSql.AutoMigrate(&models.User{})
}
