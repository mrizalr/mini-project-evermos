package mysql

import (
	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type mysqlStoreRepository struct {
	db *gorm.DB
}

func NewMysqlStoreRepository(db *gorm.DB) domain.StoreRepository {
	return &mysqlStoreRepository{db}
}

func (r *mysqlStoreRepository) GetMyStore(userID int) (domain.Store, error) {
	store := domain.Store{}
	tx := r.db.Where("user_id = ?", userID).Find(&store)
	return store, tx.Error
}

func (r *mysqlStoreRepository) GetStoreByName(storeName string) (domain.Store, error) {
	store := domain.Store{}
	tx := r.db.Where("name = ?", storeName).Find(&store)
	return store, tx.Error
}

func (r *mysqlStoreRepository) UpdateStore(store domain.Store) error {
	tx := r.db.Model(&store).Updates(&store)
	return tx.Error
}

func (r *mysqlStoreRepository) GetStoreByID(storeID int) (domain.Store, error) {
	store := domain.Store{}
	tx := r.db.Where("id = ?", storeID).First(&store)
	return store, tx.Error
}

func (r *mysqlStoreRepository) GetStores(opts model.GetStoresOptions) ([]domain.Store, error) {
	stores := []domain.Store{}
	tx := r.db.Limit(opts.Limit).Offset(opts.Limit * (opts.Page - 1))

	if opts.Name != "" {
		tx = tx.Where("name LIKE ?", "%"+opts.Name+"%")
	}

	tx.Find(&stores)
	return stores, tx.Error
}
