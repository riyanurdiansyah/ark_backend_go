package database

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupGetConnection() *gorm.DB {
	// err := godotenv.Load(".env")
	// helper.PanicIfError(err)
	dbUser := "root"
	dbHost := "localhost"
	dbName := "arkademi"

	dsn := dbUser + ":@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("gagal terhubung ke database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("gagal terhubung ke database")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	return db

}
