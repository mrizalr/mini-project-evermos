package domain

import (
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name           string          `json:"nama_produk"`
	Slug           string          `json:"slug"`
	ResellerPrice  float32         `json:"harga_reseler"`
	ConsumentPrice float32         `json:"harga_konsumen"`
	Stock          int             `json:"stok"`
	Description    string          `json:"deskripsi"`
	StoreID        uint            `json:"id_toko"`
	CategoryID     uint            `json:"id_category"`
	Store          Store           `json:"toko"`
	Category       Category        `json:"category"`
	Photos         []ProductPhotos `json:"photos"`
}

type ProductPhotos struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	Url       string `json:"url"`
}

type ProductRepository interface {
	CreateProduct(Product) (int, error)
	GetProductByID(int) (Product, error)
	GetProducts(model.GetProductOptions) ([]Product, error)
	DeleteProductByID(int) error
	UpdateProduct(Product) error
}

type ProductUsecase interface {
	CreateNewProduct(model.CreateProductRequest) (int, error)
	GetProductByID(int) (model.GetProductResponse, error)
	GetProducts(model.GetProductOptions) ([]model.GetProductResponse, error)
	DeleteProductByID(int, int) error
	UpdateProduct(int, int, model.CreateProductRequest) error
}
