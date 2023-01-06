package mysql

import (
	"github.com/mrizalr/mini-project-evermos/domain"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{db}
}

func (r *mysqlUserRepository) Register(user domain.User) error {
	tx := r.db.Create(&user)
	return tx.Error
}
