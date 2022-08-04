package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"os"

	log "github.com/sirupsen/logrus"
)

// Database is a
type Database struct{}

// ConnectDB is a
func (o Database) ConnectDB() (*sql.DB, error) {
	DB := os.Getenv("DB")
	URI := os.Getenv("DBURL")

	db, err := sql.Open(DB, URI)
	//db, err := sql.Open(DB, URI)

	if err != nil {
		log.Warnf("failed connection to DB : %v", err)
		return nil, fmt.Errorf("failed connection to DB : %v", err)
	}

	return db, nil
}
func (o Database) ConnectDBGorm() (*gorm.DB, error) {
	// DB := os.Getenv("DB")
	URI := os.Getenv("DBURL")
	dsn := URI + "?charset=utf8&parseTime=True&loc=Asia%2FJakarta"
	// sqlDB, err := sql.Open(DB, FULLURI)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	//db, err := sql.Open(DB, URI)

	if err != nil {
		log.Warnf("failed connection to DB : %v", err)
		return nil, fmt.Errorf("failed connection to DB : %v", err)
	}

	return db, nil
}
