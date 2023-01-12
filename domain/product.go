package domain

import (
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name           string  `json:"nama_produk"`
	Slug           string  `json:"slug"`
	ResellerPrice  float32 `json:"harga_reseler"`
	ConsumentPrice float32 `json:"harga_konsumen"`
	Stock          int     `json:"stok"`
	Description    string  `json:"deskripsi"`
	StoreID        uint    `json:"id_toko"`
	CategoryID     uint    `json:"id_category"`
}

type ProductPhotos struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	Url       string `json:"url"`
}

type ProductRepository interface {
	CreateProduct(Product, []ProductPhotos) (int, error)
	GetProductByID(int) (Product, []ProductPhotos, error)
	GetProducts() ([]Product, []ProductPhotos, error)
}

type ProductUsecase interface {
	CreateNewProduct(int, model.CreateProductRequest) (int, error)
	GetProductByID(int) (model.GetProductResponse, error)
	GetProducts() ([]model.GetProductResponse, error)
}
