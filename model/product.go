package model

type CreateProductRequest struct {
	Name           string   `json:"nama_produk"`
	ResellerPrice  float32  `json:"harga_reseler"`
	ConsumentPrice float32  `json:"harga_konsumen"`
	Stock          int      `json:"stok"`
	Description    string   `json:"deskripsi"`
	CategoryID     uint     `json:"id_category"`
	Photos         []string `json:"photos"`
}

type GetProductResponse struct {
	ID             int                 `json:"id"`
	Name           string              `json:"nama_produk"`
	Slug           string              `json:"slug"`
	ResellerPrice  float32             `json:"harga_reseler"`
	ConsumentPrice float32             `json:"harga_konsumen"`
	Stock          int                 `json:"stok"`
	Description    string              `json:"deskripsi"`
	Store          GetStoreResponse    `json:"toko"`
	Category       GetCategoryResponse `json:"category"`
	Photos         []struct {
		ID        uint   `json:"id"`
		ProductID uint   `json:"product_id"`
		Url       string `json:"url"`
	} `json:"photos"`
}
