package database

import (
	"TODO-API/internal/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connect() (*gorm.DB, error) {
	// MySQL 连接参数
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	dbname := "todo"

	// 构造 DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	// 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy:         schema.NamingStrategy{},
	})
	if err != nil {
		return nil, err
	}

	// 自动迁移数据库结构
	err = db.AutoMigrate(&model.Todo{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
