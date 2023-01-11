package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string    `json:"nama" gorm:"type:varchar(255);not null"`
	Password    string    `json:"kata_sandi" gorm:"not null"`
	PhoneNumber string    `json:"no_telp" gorm:"type:varchar(255);not null;unique"`
	Birthdate   time.Time `json:"tanggal_lahir"`
	Bio         string    `json:"tentang"`
	Job         string    `json:"pekerjaan"`
	Email       string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	ProvinceID  uint      `json:"id_provinsi"`
	CityID      uint      `json:"id_kota"`
	Role        string    `json:"status"`
}

func (u *User) AfterCreate(tx *gorm.DB) error {
	store := Store{
		Name:     fmt.Sprintf("toko-%s", uuid.New().String()),
		PhotoURL: "",
		UserID:   u.ID,
	}

	return tx.Create(&store).Error
}

type UserRepository interface {
	Register(User) error
	GetUserByPhoneNumber(string) (User, error)
	GetUserByID(int) (User, error)
	UpdateUser(User) error
}
type UserUsecase interface {
	Register(model.UserRegisterRequest) error
	Login(model.UserLoginRequest) (model.UserLoginResponse, error)
	GetMyProfile(int) (model.GetUserResponse, error)
	UpdateMyProfile(int, model.UpdateUserRequest) error
}
