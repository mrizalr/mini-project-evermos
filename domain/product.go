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
	Store          struct {
		ID       uint   `json:"id"`
		Name     string `json:"nama_toko"`
		PhotoURL string `json:"url_foto"`
	} `json:"toko"`
	Category model.GetCategoryResponse `json:"category"`
	Photos   []struct {
		ID        uint   `json:"id"`
		ProductID uint   `json:"product_id"`
		Url       string `json:"url"`
	} `json:"photos"`
}

type ProductRepository interface{}
type ProductUsecase interface{}
