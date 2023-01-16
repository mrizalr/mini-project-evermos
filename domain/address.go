package domain

import (
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Title        string `json:"judul_alamat" gorm:"type:varchar(50)"`
	ReceiverName string `json:"nama_penerima" gorm:"type:varchar(50)"`
	PhoneNumber  string `json:"no_telp" gorm:"type:varchar(20)"`
	Detail       string `json:"detail_alamat" gorm:"type:varchar(255)"`
	UserID       uint
}

type AddressRepository interface {
	CreateNewAddress(Address) (int, error)
	GetMyAddress(int) ([]Address, error)
	GetAddressByID(int) (Address, error)
	UpdateAddress(Address) error
	DeleteAddress(int) error
}

type AddressUsecase interface {
	CreateNewAddress(address Address) (int, error)
	GetMyAddress(int) ([]model.GetAddressResponse, error)
	GetAddressByID(int, int) (model.GetAddressResponse, error)
	UpdateAddress(int, int, model.UpdateAddressRequest) error
	DeleteAddress(int, int) error
}
