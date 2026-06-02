package database

import (
	"fmt"
	"log"

	"github.com/sachithramanamperi/croptracetelemetry/config"
	"github.com/sachithramanamperi/croptracetelemetry/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto Migrate the schema
	err = DB.AutoMigrate(&models.TelemetryRecord{})
	if err != nil {
		log.Printf("Failed to run auto migrations: %v", err)
	}

	return nil
}

func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
}
