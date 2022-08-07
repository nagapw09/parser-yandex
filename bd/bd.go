package bd

import (
	"github.com/nagapw09/parser2/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func open() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/parser?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	logger.ForError(err)
	return db
}

func Base() {

	err := open().Exec("CREATE TABLE google_queries( Id INT AUTO_INCREMENT PRIMARY KEY, link VARCHAR(220), header VARCHAR(220), request VARCHAR(220), pages VARCHAR(220) );")
	if err != nil {
		log.Println(err.Error)
	}
}

func Insert(link string, header string, request string, pages string) *gorm.DB {

	err := open().Exec("INSERT INTO `google_queries`(`link`, `header`, `request`, `pages`) VALUES (?, ?, ?, ?)",
		link, header, request, pages)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
