package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    // Replace with your actual MySQL credentials and connection details
    dsn := "root:bDDPgmKmVtAyhgxQkLzsOLQaqdFRvyDW@tcp(viaduct.proxy.rlwy.net:21337)/railway?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    DB = db
    fmt.Println("Database connected successfully")
}

func GetDB() *gorm.DB {
    return DB
}
