package usecase

import (
	"fmt"
	"strings"

	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type categoryUsecase struct {
	categoryRepository domain.CategoryRepository
}

func NewCategoryUsecase(categoryRepository domain.CategoryRepository) domain.CategoryUsecase {
	return &categoryUsecase{categoryRepository}
}

func (u *categoryUsecase) CreateNewCategory(userRole string, addCategoryRequest model.AddCategoryRequest) (int, error) {
	if strings.ToLower(userRole) != "admin" {
		return 0, fmt.Errorf("forbidden, only an admin can add a category")
	}

	category := domain.Category{
		Name: addCategoryRequest.Name,
	}

	return u.categoryRepository.CreateCategory(category)
}

func (u *categoryUsecase) GetCategories() ([]model.GetCategoryResponse, error) {
	result := []model.GetCategoryResponse{}
	categories, err := u.categoryRepository.GetCategories()
	if err != nil {
		return result, err
	}

	for _, cat := range categories {
		categoryResponse := model.GetCategoryResponse{
			ID:   cat.ID,
			Name: cat.Name,
		}
		result = append(result, categoryResponse)
	}

	return result, nil
}

func (u *categoryUsecase) GetCategoryByID(categoryID int) (model.GetCategoryResponse, error) {
	result := model.GetCategoryResponse{}
	category, err := u.categoryRepository.GetCategoryByID(categoryID)
	if err != nil {
		return result, err
	}

	result.ID = category.ID
	result.Name = category.Name

	return result, nil
}

func (u *categoryUsecase) DeleteCategoryByID(userRole string, categoryID int) error {
	if strings.ToLower(userRole) != "admin" {
		return fmt.Errorf("forbidden, only an admin can add a category")
	}

	return u.categoryRepository.DeleteCategoryByID(categoryID)
}

func (u *categoryUsecase) UpdateCategory(userRole string, categoryID int, categoryRequest model.AddCategoryRequest) error {
	if strings.ToLower(userRole) != "admin" {
		return fmt.Errorf("forbidden, only an admin can add a category")
	}

	category := domain.Category{
		Model: gorm.Model{
			ID: uint(categoryID),
		},
		Name: categoryRequest.Name,
	}

	return u.categoryRepository.UpdateCategory(category)
}
