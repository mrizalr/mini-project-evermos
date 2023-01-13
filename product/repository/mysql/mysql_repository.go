package mysql

import (
	"fmt"

	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	db *gorm.DB
}

func NewMysqlProductRepository(db *gorm.DB) domain.ProductRepository {
	return &mysqlProductRepository{db}
}

func (r *mysqlProductRepository) CreateProduct(product domain.Product) (int, error) {
	tx := r.db.Create(&product)
	return int(product.ID), tx.Error
}

func (r *mysqlProductRepository) GetProductByID(productID int) (domain.Product, error) {
	products := domain.Product{}
	tx := r.db.Where("id = ?", productID).Preload("Store").Preload("Category").Preload("Photos").Find(&products)
	return products, tx.Error
}

func (r *mysqlProductRepository) GetProducts(opts model.GetProductOptions) ([]domain.Product, error) {
	products := []domain.Product{}
	fmt.Println(opts)
	tx := r.db.Preload("Store").Preload("Category").Preload("Photos")

	if opts.Name != "" {
		tx.Where("name = ?", opts.Name)
	}

	if opts.CategoryID != 0 {
		tx.Where("category_id = ?", opts.CategoryID)
	}

	if opts.StoreID != 0 {
		tx.Where("store_id = ?", opts.StoreID)
	}

	if opts.MaxPrice != 0 {
		tx.Where("consument_price <= ?", opts.MaxPrice)
	}

	if opts.MinPrice != 0 {
		tx.Where("consument_price >= ?", opts.MinPrice)
	}

	tx.Find(&products)
	return products, tx.Error
}

func (r *mysqlProductRepository) DeleteProductByID(productID int) error {
	tx := r.db.Where("id = ?", productID).Delete(&domain.Product{})
	if tx.Error != nil {
		return tx.Error
	}

	tx = r.db.Where("product_id = ?", productID).Delete(&domain.ProductPhotos{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *mysqlProductRepository) UpdateProduct(product domain.Product) error {
	tx := r.db.Where("product_id = ?", product.ID).Delete(&domain.ProductPhotos{})
	if tx.Error != nil {
		return tx.Error
	}

	tx = r.db.Model(&product).Updates(&product)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
