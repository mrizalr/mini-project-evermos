package usecase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/model"
	"github.com/mrizalr/mini-project-evermos/utils"
	"github.com/spf13/viper"
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
		Bio:         "",
		Job:         request.Job,
		Email:       request.Email,
		ProvinceID:  uint(userProvinceID),
		CityID:      uint(userCityID),
		Role:        "user",
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return u.userRepository.Register(user)
}

func (u *userUsecase) Login(request model.UserLoginRequest) (model.UserLoginResponse, error) {
	var userResponse model.UserLoginResponse

	user, err := u.userRepository.GetUserByPhoneNumber(request.PhoneNumber)
	if err != nil {
		return userResponse, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return userResponse, err
	}

	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/province/%d.json", user.ProvinceID)
	userProvince := model.Province{}
	err = utils.GetRegionData(url, &userProvince)
	if err != nil {
		return userResponse, err
	}

	url = fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regency/%d.json", user.CityID)
	userCity := model.City{}
	err = utils.GetRegionData(url, &userCity)
	if err != nil {
		return userResponse, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   strconv.Itoa(int(user.ID)),
		"user_role": user.Role,
		"exp":       strconv.FormatInt(time.Now().Add(24*time.Hour).Unix(), 10),
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("secret_key")))
	if err != nil {
		return userResponse, err
	}

	userResponse = model.UserLoginResponse{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		BirthDate:   user.Birthdate.Format("02/01/2006"),
		Bio:         user.Bio,
		Job:         user.Job,
		Email:       user.Email,
		ProvinceID:  userProvince,
		CityID:      userCity,
		Token:       tokenString,
	}

	return userResponse, nil
}

func (u *userUsecase) GetMyProfile(userID int) (model.GetUserResponse, error) {
	var userResponse model.GetUserResponse

	user, err := u.userRepository.GetUserByID(userID)
	if err != nil {
		return userResponse, err
	}

	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/province/%d.json", user.ProvinceID)
	userProvince := model.Province{}
	err = utils.GetRegionData(url, &userProvince)
	if err != nil {
		return userResponse, err
	}

	url = fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regency/%d.json", user.CityID)
	userCity := model.City{}
	err = utils.GetRegionData(url, &userCity)
	if err != nil {
		return userResponse, err
	}

	userResponse = model.GetUserResponse{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		BirthDate:   user.Birthdate.Format("02/01/2006"),
		Bio:         user.Bio,
		Job:         user.Job,
		Email:       user.Email,
		ProvinceID:  userProvince,
		CityID:      userCity,
	}

	return userResponse, nil
}

func (u *userUsecase) UpdateMyProfile(userID int, updateUserRequest model.UpdateUserRequest) error {
	userBirthdate, err := time.Parse("02/01/2006", updateUserRequest.Birthdate)
	if err != nil {
		return err
	}

	userProvinceID, err := strconv.Atoi(updateUserRequest.ProvinceID)
	if err != nil {
		return err
	}

	userCityID, err := strconv.Atoi(updateUserRequest.CityID)
	if err != nil {
		return err
	}

	user := domain.User{
		Model:       gorm.Model{ID: uint(userID)},
		Name:        updateUserRequest.Name,
		Password:    updateUserRequest.Password,
		PhoneNumber: updateUserRequest.PhoneNumber,
		Birthdate:   userBirthdate,
		Bio:         updateUserRequest.Bio,
		Job:         updateUserRequest.Job,
		Email:       updateUserRequest.Email,
		ProvinceID:  uint(userProvinceID),
		CityID:      uint(userCityID),
	}

	if updateUserRequest.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(updateUserRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user.Password = string(hashed)
	}

	return u.userRepository.UpdateUser(user)
}
