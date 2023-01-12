package mysql

import (
	"github.com/mrizalr/mini-project-evermos/domain"
	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	db *gorm.DB
}

func NewMysqlProductRepository(db *gorm.DB) domain.ProductRepository {
	return &mysqlProductRepository{db}
}

func (r *mysqlProductRepository) withTransaction(fn func(tx *gorm.DB) error) error {
	tx := r.db.Begin()
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *mysqlProductRepository) CreateProduct(product domain.Product, photos []domain.ProductPhotos) (int, error) {
	err := r.withTransaction(func(tx *gorm.DB) error {
		if err := tx.Create(&product).Error; err != nil {
			return err
		}

		for _, p := range photos {
			p.ProductID = product.ID
			if err := tx.Create(&p).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return int(product.ID), err
}

func (r *mysqlProductRepository) GetProductByID(productID int) (domain.Product, []domain.ProductPhotos, error) {
	product := domain.Product{}
	productPhotos := []domain.ProductPhotos{}

	err := r.db.Where("id = ?", productID).First(&product).Error
	if err != nil {
		return product, productPhotos, err
	}

	err = r.db.Where("product_id = ?", productID).Find(&productPhotos).Error
	if err != nil {
		return product, productPhotos, err
	}

	return product, productPhotos, nil
}

func (r *mysqlProductRepository) GetProducts() ([]domain.Product, []domain.ProductPhotos, error) {
	products := []domain.Product{}
	productPhotos := []domain.ProductPhotos{}

	err := r.db.Find(&products).Error
	if err != nil {
		return products, productPhotos, err
	}

	err = r.db.Find(&productPhotos).Error
	if err != nil {
		return products, productPhotos, err
	}

	return products, productPhotos, nil
}
