package bd

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func open() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/parser?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func Base() *gorm.DB {

	err := open().Exec("CREATE TABLE google_queries( Id INT AUTO_INCREMENT PRIMARY KEY, link VARCHAR(220), header VARCHAR(220), pages VARCHAR(220) );")

	if err != nil {
		return err
	}
	return nil
}

func Insert(link string, header string, pages string) *gorm.DB {

	err := open().Exec("INSERT INTO `google_queries`(`link`, `header`, `pages`) VALUES (?, ?, ?)",
		link, header, pages)
	if err != nil {
		return err
	}
	return nil
}
