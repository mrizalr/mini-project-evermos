package usecase

import "github.com/mrizalr/mini-project-evermos/domain"

type productUsecase struct {
	productRepository domain.ProductRepository
}

func NewProductUsecase(productRepository domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{productRepository}
}
