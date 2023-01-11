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

func (r *mysqlUserRepository) GetUserByPhoneNumber(phoneNumber string) (domain.User, error) {
	user := domain.User{}
	tx := r.db.Where("phone_number = ?", phoneNumber).First(&user)
	return user, tx.Error
}

func (r *mysqlUserRepository) GetUserByID(userID int) (domain.User, error) {
	user := domain.User{}
	tx := r.db.Where("id = ?", userID).First(&user)
	return user, tx.Error
}

func (r *mysqlUserRepository) UpdateUser(user domain.User) error {
	tx := r.db.Model(&user).Updates(&user)
	return tx.Error
}
