package mysql

import (
	"github.com/mrizalr/mini-project-evermos/domain"
	"gorm.io/gorm"
)

type mysqlAddressRepository struct {
	db *gorm.DB
}

func NewMysqlAddressRepository(db *gorm.DB) domain.AddressRepository {
	return &mysqlAddressRepository{db}
}

func (r *mysqlAddressRepository) CreateNewAddress(address domain.Address) (int, error) {
	tx := r.db.Create(&address)
	return int(address.ID), tx.Error
}

func (r *mysqlAddressRepository) GetMyAddress(userID int) ([]domain.Address, error) {
	addresses := []domain.Address{}
	tx := r.db.Where("user_id = ?", userID).Find(&addresses)
	return addresses, tx.Error
}

func (r *mysqlAddressRepository) GetAddressByID(addressID int) (domain.Address, error) {
	address := domain.Address{}
	tx := r.db.Where("id = ?", addressID).Find(&address)
	return address, tx.Error
}

func (r *mysqlAddressRepository) UpdateAddress(address domain.Address) error {
	tx := r.db.Model(&address).Updates(&address)
	return tx.Error
}

func (r *mysqlAddressRepository) DeleteAddress(addressID int) error {
	tx := r.db.Where("id = ?", addressID).Delete(&domain.Address{})
	return tx.Error
}
