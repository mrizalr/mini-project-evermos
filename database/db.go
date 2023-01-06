package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", viper.GetString("db.user"), viper.GetString("db.password"), viper.GetString("db.port"), viper.GetString("db.dbname"))
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
