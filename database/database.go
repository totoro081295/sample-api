package database

import (
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

// ConnectDB データベース接続
func ConnectDB() *gorm.DB {
	dbURI := os.Getenv("DATABASE_URL")
	logLevelInt, _ := strconv.Atoi(os.Getenv("LOGLEVEL"))
	logLevel := logger.LogLevel(logLevelInt)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		panic(err)
	}

	return db
}

func Close(odb *gorm.DB) {
	if sqlDB, err := odb.DB(); err != nil {
		panic(err)
	} else {
		if err := sqlDB.Close(); err != nil {
			panic(err)
		}
	}
}
