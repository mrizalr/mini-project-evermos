package domain

import (
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	Name     string `json:"nama_toko" gorm:"type:varchar(150)"`
	PhotoURL string `json:"url_foto"`
	UserID   uint   `json:"user_id"`
}

type StoreRepository interface {
	GetMyStore(int) (Store, error)
	GetStoreByName(string) (Store, error)
	UpdateStore(Store) error
	GetStoreByID(int) (Store, error)
	GetStores(model.GetStoresOptions) ([]Store, error)
}

type StoreUsecase interface {
	GetMyStore(int) (Store, error)
	UpdateStore(int, int, model.UpdateStoreRequest) error
	GetStoreByID(int) (Store, error)
	GetStores(model.GetStoresOptions) ([]Store, error)
}
