package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type Book struct {
	ID          int
	Title       string
	Description string
	PageCount   int
	PubYear     int
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "bohdan:Bohdan.2525@tcp(127.0.0.1:3306)/shop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetConnMaxLifetime(2 * time.Second)

	book := &Book{}
	res := db.First(book, 3)
	if res.Error != nil {
		log.Fatal(res.Error)
	}
	fmt.Print(book)
}
