package database

import (
	"fmt"
	"go-api/pkg"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	fmt.Println("Connecting to database...")
	var host string = pkg.GetEnv("DB_HOST", "localhost")
	var user string = pkg.GetEnv("DB_USER", "postgres")
	var password string = pkg.GetEnv("DB_PASSWORD", "postgres")
	var dbname string = pkg.GetEnv("DB_NAME", "postgres")
	var port string = pkg.GetEnv("DB_PORT", "5432")

	connectionStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionStr,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db
	fmt.Println("Connected to database")
	return db
}
