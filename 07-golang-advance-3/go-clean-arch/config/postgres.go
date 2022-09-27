package config

import (
	_ "github.com/joho/godotenv/autoload"
	"go-clean-arch/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var err error

func Database() {
	//dsn := "host=localhost user= dbname= port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	//DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB, err = gorm.Open(postgres.Open(os.Getenv("db_url")), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	log.Println("Database Connected")
}

func Migrate() {
	err := DB.AutoMigrate(&entity.User{})
	if err != nil {
		return
	}
}
