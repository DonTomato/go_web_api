package db

import (
	"fmt"

	"github.com/dontomato/go_web_api/global"
	"github.com/jinzhu/gorm"
)

// Create creates gorm-db connection
func Create() (*gorm.DB, error) {
	cn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", global.Config.Host, global.Config.DBPort, global.Config.User, global.Config.Database, global.Config.Password)
	fmt.Println(cn)
	return gorm.Open("postgres", cn)
}

func test1() {
	db, _ := Create()
	db.Close()
}

// CheckDbVersion checks db version
func CheckDbVersion() bool {
	db, err := Create()
	if err != nil {
		panic(err.Error)
	}

	if !db.HasTable(&Product{}) {
		db.CreateTable(&Product{})
	}
	return true
}
