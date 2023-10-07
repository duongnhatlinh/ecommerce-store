package database

import (
	"ecommercestore/internal/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectionDatabase(conf conf.Config) {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DbUser,
		conf.DbPassword,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("connect to database successfully")
}
