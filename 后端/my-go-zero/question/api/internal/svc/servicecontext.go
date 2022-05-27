package svc

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"question/api/internal/config"
	"question/api/internal/models"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 连接数据库
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
	)

	db, err := gorm.Open("postgres", dsn)

	db.SingularTable(true)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("连接成功")
	}

	// 自动同步更新表结构
	db.AutoMigrate(&models.QuestionBank{})

	return &ServiceContext{
		Config:  c,
		DbEngin: db,
	}
}
