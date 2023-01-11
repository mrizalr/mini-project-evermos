package mysql

import (
	"github.com/mrizalr/mini-project-evermos/domain"
	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	db *gorm.DB
}

func NewMysqlProductRepository(db *gorm.DB) domain.ProductRepository {
	return &mysqlProductRepository{db}
}
