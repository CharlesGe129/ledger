package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"ledger/pkg/models"
	"sync"
)

var db *gorm.DB
var err error

func GetDB() *gorm.DB {
	var once sync.Once
	once.Do(func() {
		dbStr := "root:toor123@tcp(127.0.0.1:33600)/ledger?charset=utf8&parseTime=True&loc=Local"
		db, err = gorm.Open("mysql", dbStr)
		if err != nil {
			panic(err)
		}

		doMigration(db)
	})
	return db
}

func doMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Item{})
}
