package db

import "github.com/jinzhu/gorm"

func DbInit(driver, url string) (*gorm.DB, error) {
	return gorm.Open(driver, url)
}
