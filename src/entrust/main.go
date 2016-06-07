package main

import (
	logger "github.com/donnie4w/go-logger/logger"
	"config"
	//"fmt"
)

func main() {
	//defer logger.Flush()
	logger.SetRollingFile("log", "entrust.log", 10, 5, logger.KB)
	logger.SetRollingDaily("log", "entrust.log")
	logger.SetLevel(logger.INFO)
	//fmt.Println("Hello, GO!")
	logger.Info("hello go!")
	config.LoadConfig()
}