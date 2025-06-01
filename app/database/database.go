package database

import (
	"fmt"
	"log"
	"time"
	"time-doo-api/app/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(cfg *config.Database) *gorm.DB {
	db, err := gorm.Open(postgres.Open(getDatabaseDsn(cfg)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("cannot connect to DB: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("cannot get db instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.Exec(fmt.Sprintf("SET search_path TO %s", cfg.Path))

	fmt.Println("Connected to Database")
	return db
}

func getDatabaseDsn(cfg *config.Database) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s search_path=%s port=%s sslmode=%s TimeZone=Asia/Bangkok",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Path, cfg.Port, cfg.SSLMode,
	)
}
