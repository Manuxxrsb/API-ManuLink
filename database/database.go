package database

import (
	"api-manulink/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenGormDB() (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(config.DBURL()), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
    })
    if err != nil {
        return nil, err
    }

    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }

    // Configura el tiempo máximo de vida de las conexiones
    sqlDB.SetConnMaxLifetime(1 *time.Hour)

    return db, nil
}