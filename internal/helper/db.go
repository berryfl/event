package helper

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

const (
	dsn = "root:bRsSsB8e@unix(/tmp/mysql.sock)/event_db?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	db *gorm.DB
)

func InitDB() error {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("open mysql connection failed: %w", err)
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
