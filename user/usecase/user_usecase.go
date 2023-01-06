package usecase

import (
	"strconv"
	"time"

	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepository}
}

func (u *userUsecase) Register(request model.UserRegisterRequest) error {
	userBirthdate, err := time.Parse("02/01/2006", request.Birthdate)
	if err != nil {
		return err
	}

	userProvinceID, err := strconv.Atoi(request.ProvinceID)
	if err != nil {
		return err
	}

	userCityID, err := strconv.Atoi(request.CityID)
	if err != nil {
		return err
	}

	user := domain.User{
		Model:       gorm.Model{},
		Name:        request.Name,
		Password:    request.Password,
		PhoneNumber: request.PhoneNumber,
		Birthdate:   userBirthdate,
		Job:         request.Job,
		Email:       request.Email,
		ProvinceID:  uint(userProvinceID),
		CityID:      uint(userCityID),
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return u.userRepository.Register(user)
}
