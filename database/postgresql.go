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
	var dsn string = "host=" + os.Getenv("HOSTDB") + " user=" + os.Getenv("USERDB") + " password='" + os.Getenv("PASSWORDDB") + "' dbname=" + os.Getenv("DB") + " port=" + os.Getenv("PORTDB") + " sslmode=disable TimeZone=Asia/Shanghai"

	DBPostgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		utils.ErrorDatabase("PostgreSql", err)
	}

	//db.AutoMigrate(&models.User{})
}
