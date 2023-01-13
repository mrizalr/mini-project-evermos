package usecase

import (
	"fmt"

	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type storeUsecase struct {
	storeRepository domain.StoreRepository
}

func NewStoreUsecase(storeRepository domain.StoreRepository) domain.StoreUsecase {
	return &storeUsecase{storeRepository}
}

func (u *storeUsecase) GetMyStore(userID int) (domain.Store, error) {
	return u.storeRepository.GetMyStore(userID)
}

func (u *storeUsecase) UpdateStore(userID int, storeID int, updateStoreRequest model.UpdateStoreRequest) error {
	store, err := u.GetStoreByID(storeID)
	if err != nil {
		return err
	}

	if store.UserID != uint(userID) {
		return fmt.Errorf("permission denied. You are only allowed to access your own store")
	}

	foundStore, err := u.storeRepository.GetStoreByName(updateStoreRequest.Name)
	if err != nil {
		return err
	}

	emptyStore := domain.Store{}
	if foundStore != emptyStore && foundStore.ID != uint(storeID) {
		return fmt.Errorf("store name already used")
	}

	storeUpdate := domain.Store{
		Model: gorm.Model{
			ID: uint(storeID),
		},
		Name:     updateStoreRequest.Name,
		PhotoURL: updateStoreRequest.PhotoURL,
	}

	return u.storeRepository.UpdateStore(storeUpdate)

}

func (u *storeUsecase) GetStoreByID(storeID int) (domain.Store, error) {
	return u.storeRepository.GetStoreByID(storeID)
}

func (u *storeUsecase) GetStores(opts model.GetStoresOptions) ([]domain.Store, error) {
	return u.storeRepository.GetStores(opts)
}
