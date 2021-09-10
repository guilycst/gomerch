package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ref *gorm.DB

func GetConnection() *gorm.DB {
	if ref != nil {
		return ref
	}
	connStr := os.Getenv("POSTGRES_CONNECTION_STRING")
	gdb, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return gdb
}
