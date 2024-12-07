package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// InitDB initializes the database connection once and returns the singleton instance.
func InitDB() *gorm.DB {
	once.Do(func() {
		// Load .env file
		if err := godotenv.Load(); err != nil {
			log.Println("Error loading .env file, falling back to system environment variables")
		}

		// Load database credentials from environment variables
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		// Build the DSN (Data Source Name)
		dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

		// Connect to the database
		var openErr error
		db, openErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if openErr != nil {
			log.Fatalf("Failed to connect to the database: %v", openErr)
		}

		log.Println("Database connected successfully")
	})
	return db
}
