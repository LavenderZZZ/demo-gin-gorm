package model

import (
	"demo1/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func Database(conn string) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: newLogger,
	})

	if conn == "" || err != nil {
		util.Log().Error("mysql lost :%v", err)
		panic(err)
	}
	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	DB = db

	//自动迁移schema，保持schema 是最新的
	_ = DB.AutoMigrate(&User{})
}
