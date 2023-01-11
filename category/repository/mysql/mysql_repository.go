package mysql

import (
	"github.com/mrizalr/mini-project-evermos/domain"
	"gorm.io/gorm"
)

type mysqlCategoryRepository struct {
	db *gorm.DB
}

func NewMysqlCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &mysqlCategoryRepository{db}
}

func (r *mysqlCategoryRepository) CreateCategory(category domain.Category) (int, error) {
	tx := r.db.Create(&category)
	return int(category.ID), tx.Error
}

func (r *mysqlCategoryRepository) GetCategories() ([]domain.Category, error) {
	categories := []domain.Category{}
	tx := r.db.Find(&categories)
	return categories, tx.Error
}

func (r *mysqlCategoryRepository) GetCategoryByID(categoryID int) (domain.Category, error) {
	category := domain.Category{}
	tx := r.db.Where("id = ?", categoryID).First(&category)
	return category, tx.Error
}

func (r *mysqlCategoryRepository) DeleteCategoryByID(categoryID int) error {
	tx := r.db.Where("id = ?", categoryID).Delete(&domain.Category{})
	return tx.Error
}

func (r *mysqlCategoryRepository) UpdateCategory(category domain.Category) error {
	tx := r.db.Model(&category).Updates(&category)
	return tx.Error
}
