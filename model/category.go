package model

type AddCategoryRequest struct {
	Name string `json:"nama_category"`
}

type GetCategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"nama_category"`
}
