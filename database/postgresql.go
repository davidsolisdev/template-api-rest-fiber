package database

import (
	"os"

	"github.com/davidsolisdev/template-api-rest-fiber/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	//"github.com/davidalarcon68/ApiRest/models"
)

var DBPostgres *gorm.DB

func ConnectPostqreSql() {
	var err error
	var dsn string = "host=" + os.Getenv("HOST_DB") + " user=" + os.Getenv("USER_DB") + " password='" + os.Getenv("PASSWORD_DB") + "' dbname=" + os.Getenv("DB") + " port=" + os.Getenv("PORT_DB") + " sslmode=disable TimeZone=America/Guatemala"
	DBPostgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		utils.ErrorDatabase("PostgreSql", err)
	}

	//DBPostgres.AutoMigrate(&models.User{})
}
