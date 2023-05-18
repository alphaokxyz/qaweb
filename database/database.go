package database

import (
	"fmt"
	"qaweb/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

func Initdb() {
	dsn := config.Dbuser + ":" + config.Dbpassword + "@tcp(" + config.Dbhost + ":" + config.Dbport + ")/" + config.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Print("连接数据库失败")
	} else {
		fmt.Print("连接数据库成功")
	}
	sqlDB, _ := db.DB()

	db.AutoMigrate(&User{})

	db.AutoMigrate(&Question{})

	db.AutoMigrate(&Answer{})

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
