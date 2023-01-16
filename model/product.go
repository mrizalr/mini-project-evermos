package model

type CreateProductRequest struct {
	Name           string   `json:"nama_produk"`
	ResellerPrice  float32  `json:"harga_reseler"`
	ConsumentPrice float32  `json:"harga_konsumen"`
	Stock          int      `json:"stok"`
	Description    string   `json:"deskripsi"`
	CategoryID     uint     `json:"id_category"`
	StoreID        uint     `json:"id_toko"`
	Photos         []string `json:"photos"`
}

type ProductPhotosResponse struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	Url       string `json:"url"`
}

type GetProductTrxResponse struct {
	ID             int                     `json:"id"`
	Name           string                  `json:"nama_produk"`
	Slug           string                  `json:"slug"`
	ResellerPrice  float32                 `json:"harga_reseler"`
	ConsumentPrice float32                 `json:"harga_konsumen"`
	Description    string                  `json:"deskripsi"`
	Store          GetStoreTrxResponse     `json:"toko"`
	Category       GetCategoryResponse     `json:"category"`
	Photos         []ProductPhotosResponse `json:"photos"`
}

type GetProductResponse struct {
	ID             int                     `json:"id"`
	Name           string                  `json:"nama_produk"`
	Slug           string                  `json:"slug"`
	ResellerPrice  float32                 `json:"harga_reseler"`
	ConsumentPrice float32                 `json:"harga_konsumen"`
	Stock          int                     `json:"stok"`
	Description    string                  `json:"deskripsi"`
	Store          GetStoreResponse        `json:"toko"`
	Category       GetCategoryResponse     `json:"category"`
	Photos         []ProductPhotosResponse `json:"photos"`
}

type GetProductOptions struct {
	Name       string  `query:"nama_produk"`
	CategoryID int     `query:"category_id"`
	StoreID    int     `query:"toko_id"`
	MaxPrice   float32 `query:"max_harga"`
	MinPrice   float32 `query:"min_harga"`
}
