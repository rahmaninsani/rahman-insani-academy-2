package config

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"go-echo-clean-arch/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var err error

func Database() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"), os.Getenv("DB_TIMEZONE"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
