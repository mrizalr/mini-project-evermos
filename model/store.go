package model

type UpdateStoreRequest struct {
	Name     string `json:"nama_toko"`
	PhotoURL string `json:"url_foto"`
}

type GetStoresOptions struct {
	Page  int
	Limit int
	Nama  string
}
