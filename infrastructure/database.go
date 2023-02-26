package infrastructure

import (
	"go-app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func init() {
	config.New()
}

var DB *gorm.DB

func DatabaseConnect() *gorm.DB {
	dns := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USERNAME") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"
	//dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable ", host, username, password, dbName)

	result, err := gorm.Open(postgres.Open(dns))

	if err != nil {
		log.Print("db error: ", err)
		panic(err)
	}
	DB = result
	return DB
}
