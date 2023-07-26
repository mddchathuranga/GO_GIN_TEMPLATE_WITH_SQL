package db

import (
	"database/sql"

	"github.com/user/test_template/logger"
	"github.com/user/test_template/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const dsn = "root:Bb_1188484@tcp(127.0.0.1:3306)/"
const dbname = "usergo"
const dbinfo = "?charset=utf8&parseTime=True&loc=Local"

func createDatabaseIfNotExists(db *sql.DB) error {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	return err
}

func InitialDbConnection() {
	logger := logger.GetLogger() // Get the initialized logger instance
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Info("Cannot connect to DB:", err.Error())
		panic("cannot connect to the database")
	}
	defer db.Close()

	if err := createDatabaseIfNotExists(db); err != nil {
		logger.Info("Cannot create DB:", err.Error())
		panic("cannot create the database")
	}

	DB, err = gorm.Open(mysql.Open(dsn+dbname+dbinfo), &gorm.Config{})
	if err != nil {

		logger.Info("Cannot connect to DB :", err.Error())
		panic("cabbot connect to database")
	}
	DB.AutoMigrate(&models.User{})
}
