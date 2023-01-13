package usecase

import (
	"fmt"

	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type addressUsecase struct {
	addressRepository domain.AddressRepository
}

func NewAddressUsecase(addressRepository domain.AddressRepository) domain.AddressUsecase {
	return &addressUsecase{addressRepository}
}

func (u *addressUsecase) CreateNewAddress(address domain.Address) (int, error) {
	return u.addressRepository.CreateNewAddress(address)
}

func (u *addressUsecase) GetMyAddress(userID int) ([]model.GetAddressResponse, error) {
	result := []model.GetAddressResponse{}
	addresses, err := u.addressRepository.GetMyAddress(userID)
	if err != nil {
		return result, err
	}

	for _, address := range addresses {
		addressResponse := model.GetAddressResponse{
			ID:           address.ID,
			Title:        address.Title,
			ReceiverName: address.ReceiverName,
			PhoneNumber:  address.PhoneNumber,
			Detail:       address.Detail,
		}

		result = append(result, addressResponse)
	}

	return result, nil
}

func (u *addressUsecase) GetAddressByID(userID int, addressID int) (model.GetAddressResponse, error) {
	var result model.GetAddressResponse
	address, err := u.addressRepository.GetAddressByID(addressID)
	if err != nil {
		return result, err
	}

	if address.UserID != uint(userID) {
		return result, fmt.Errorf("permission denied. You are only allowed to access your own personal information")
	}

	result = model.GetAddressResponse{
		ID:           address.ID,
		Title:        address.Title,
		ReceiverName: address.ReceiverName,
		PhoneNumber:  address.PhoneNumber,
		Detail:       address.Detail,
	}

	return result, nil
}

func (u *addressUsecase) UpdateAddress(userID int, addressID int, updateAddressRequest model.UpdateAddressRequest) error {
	address, err := u.addressRepository.GetAddressByID(addressID)
	if err != nil {
		return err
	}

	if address.UserID != uint(userID) {
		return fmt.Errorf("permission denied. You are only allowed to access your own personal information")
	}

	updateAddress := domain.Address{
		Model: gorm.Model{
			ID: uint(addressID),
		},
		ReceiverName: updateAddressRequest.ReceiverName,
		PhoneNumber:  updateAddressRequest.PhoneNumber,
		Detail:       updateAddressRequest.Detail,
	}

	return u.addressRepository.UpdateAddress(updateAddress)
}

func (u *addressUsecase) DeleteAddress(userID int, addressID int) error {
	address, err := u.addressRepository.GetAddressByID(addressID)
	if err != nil {
		return err
	}

	if address.UserID != uint(userID) {
		return fmt.Errorf("permission denied. You are only allowed to access your own personal information")
	}

	return u.addressRepository.DeleteAddress(addressID)
}
