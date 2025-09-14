package db

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresConnection() (*gorm.DB, error) {
	driver := os.Getenv("DB_DRIVER")
	if driver != "postgres" {
		return nil, fmt.Errorf("unsuported driver: %s", driver)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslmode := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", 
						host,   user,   pass,       dbName,   port,   sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	maxConnStr := os.Getenv("MAX_CONN")
	maxConn, err := strconv.Atoi(maxConnStr)
	if err != nil {
		maxConn = 50
	}

	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetMaxIdleConns(maxConn / 2)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func RunAutoMigrate(db *gorm.DB, models ...interface{}) error {
	return db.AutoMigrate(models...)
}