package db

import (
	"douyin/v1/pkg/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt: true,
			// SkipDefaultTransaction: true, 这里要开启事务
		},
	)
	println("数据库连接成功!")
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if !m.HasTable(&Video{}) {
		if err = m.CreateTable(&Video{}); err != nil {
			panic(err)
		}
	}

	if !m.HasTable(&Favorite{}) {
		if err = m.CreateTable(&Favorite{}); err != nil {
			panic(err)
		}
	}

	if !m.HasTable(&Comment{}) {
		if err = m.CreateTable(&Comment{}); err != nil {
			panic(err)
		}
	}

}
