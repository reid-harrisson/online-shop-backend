package db

import (
	"OnlineStoreBackend/pkgs/config"
	"fmt"
	"io"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(cfg *config.Config) *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&sql_mode='STRICT_ALL_TABLES,NO_ZERO_DATE,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO'",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name)

	fmt.Println(dataSourceName)

	var multipleWriter io.Writer
	if cfg.Log.DB.Path != "" {
		// Open a log file
		logFile, err := os.OpenFile(cfg.Log.DB.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		// Write log both in file and console
		multipleWriter = io.MultiWriter(logFile, os.Stdout)
	} else {
		// Write log in console
		multipleWriter = io.MultiWriter(os.Stdout)
	}

	multipleLog := log.New(multipleWriter, "", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: logger.New(multipleLog, logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		}),
	})
	if err != nil {
		panic(err.Error())
	}

	return db
}
