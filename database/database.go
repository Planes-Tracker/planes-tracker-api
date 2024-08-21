package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/LockBlock-dev/planes-tracker-api/config"
)

func ConnectDB(verbose bool) *gorm.DB {
	var dbLogger logger.Interface

	if verbose {
		dbLogger = logger.Default
	} else {
		dbLogger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(
		postgres.Open(config.DatabaseDSN()),
		&gorm.Config{
			Logger:         dbLogger,
			TranslateError: true,
		},
	)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect to the database: %w", err))
	}

	fmt.Println("Connected to the database")

	return db
}
