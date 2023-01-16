package domain

import (
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"nama_category" gorm:"type:varchar(50)"`
}

type CategoryRepository interface {
	CreateCategory(category Category) (int, error)
	GetCategories() ([]Category, error)
	GetCategoryByID(int) (Category, error)
	DeleteCategoryByID(int) error
	UpdateCategory(Category) error
}
type CategoryUsecase interface {
	CreateNewCategory(string, model.AddCategoryRequest) (int, error)
	GetCategories() ([]model.GetCategoryResponse, error)
	GetCategoryByID(int) (model.GetCategoryResponse, error)
	DeleteCategoryByID(string, int) error
	UpdateCategory(string, int, model.AddCategoryRequest) error
}
