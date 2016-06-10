package main

import (
	//logger "github.com/donnie4w/go-logger/logger"
	"config"
	"fmt"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
  gorm.Model
  Name string
}

func main() {
	//defer logger.Flush()
	//logger.SetRollingFile("log", "entrust.log", 10, 5, logger.KB)
	//logger.SetRollingDaily("log", "entrust.log")
	//logger.SetLevel(logger.INFO)
	
	db, err := gorm.Open("mysql", "root:0709@/shudang?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "jinzhu"})
	var user User
	db.first(&user)
	fmt.Println("%+v", user)
}