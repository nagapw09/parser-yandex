package bd

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func open() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/parser?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

type GoogleQuery struct {
	Id      int
	Link    string
	Header  string
	Request string
	Page    int
}

func CreateTable() error {
	connection, err := open()
	if err != nil {
		return err
	}
	err = connection.AutoMigrate(&GoogleQuery{})
	if err != nil {
		return err
	}

	return nil
}

func Insert(link string, header string, request string, page int) error {

	connection, err := open()
	if err != nil {
		return err
	}

	tx := connection.Create(&GoogleQuery{Link: link, Header: header, Request: request, Page: page})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
