package sql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB = nil

const (
	username = "root"
	password = "password"
	hostname = "127.0.0.1:3306"
	dbname   = "db"
)

func GetConnection() *gorm.DB {
	if dbConn != nil {
		return dbConn
	}

	db, err := gorm.Open(mysql.Open(dsn(dbname)), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logger.ERROR.Panicf("Connection refused .. %v", err)
	}

	// "Important settings"
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxIdleTime(time.Minute * 1)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	pingConnection(sqlDB)

	dbConn = db
	return dbConn
}

func pingConnection(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		logger.ERROR.Panic(err)
	}

	logger.INFO.Println("DB connection successful !!")
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
