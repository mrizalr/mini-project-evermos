package database

import (
	"fmt"

	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", viper.GetString("db.user"), viper.GetString("db.password"), viper.GetString("db.port"), viper.GetString("db.dbname"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("cannot connect to database : %v", err.Error()))
	}

	db.AutoMigrate(
		&domain.User{},
		&domain.Store{},
		&domain.Category{},
		&domain.Product{},
		&domain.ProductPhotos{},
		&domain.Address{},
	)
	return db
}
