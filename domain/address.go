package domain

import (
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Title        string `json:"judul_alamat"`
	ReceiverName string `json:"nama_penerima"`
	PhoneNumber  string `json:"no_telp"`
	Detail       string `json:"detail_alamat"`
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
