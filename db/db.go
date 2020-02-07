package db

import (
	"os"

	"github.com/jinzhu/gorm"
)

//
// Open database connection specific to the environment
//
func Open() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", os.Getenv("DB_USER")+"@"+os.Getenv("DB_HOST")+"/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, nil
}
