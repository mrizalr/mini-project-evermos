package domain

import (
	"time"

	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string    `json:"nama" gorm:"type:varchar(255);not null"`
	Password    string    `json:"kata_sandi" gorm:"not null"`
	PhoneNumber string    `json:"no_telp" gorm:"type:varchar(255);not null;unique"`
	Birthdate   time.Time `json:"tanggal_lahir"`
	Job         string    `json:"pekerjaan"`
	Email       string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	ProvinceID  uint      `json:"id_provinsi"`
	CityID      uint      `json:"id_kota"`
}

type UserRepository interface {
	Register(User) error
	Login(string) (User, error)
}
type UserUsecase interface {
	Register(model.UserRegisterRequest) error
	Login(model.UserLoginRequest) (model.UserLoginResponse, error)
}
