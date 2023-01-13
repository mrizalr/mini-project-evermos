package usecase

import (
	"fmt"
	"os"
	"strings"

	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type productUsecase struct {
	productRepository  domain.ProductRepository
	storeRepository    domain.StoreRepository
	categoryRepository domain.CategoryRepository
}

func NewProductUsecase(productRepository domain.ProductRepository, storeRepository domain.StoreRepository, categoryRepository domain.CategoryRepository) domain.ProductUsecase {
	return &productUsecase{
		productRepository:  productRepository,
		storeRepository:    storeRepository,
		categoryRepository: categoryRepository,
	}
}

func (u *productUsecase) CreateNewProduct(createProductRequest model.CreateProductRequest) (int, error) {
	productSlug := strings.ToLower(strings.ReplaceAll(createProductRequest.Name, " ", "-"))

	product := domain.Product{
		Name:           createProductRequest.Name,
		Slug:           productSlug,
		ResellerPrice:  createProductRequest.ResellerPrice,
		ConsumentPrice: createProductRequest.ConsumentPrice,
		Stock:          createProductRequest.Stock,
		Description:    createProductRequest.Description,
		StoreID:        createProductRequest.StoreID,
		CategoryID:     createProductRequest.CategoryID,
	}

	for _, productPhoto := range createProductRequest.Photos {
		product.Photos = append(product.Photos, domain.ProductPhotos{
			Url: productPhoto,
		})
	}

	return u.productRepository.CreateProduct(product)
}

func (u *productUsecase) GetProductByID(productID int) (model.GetProductResponse, error) {
	var result model.GetProductResponse

	product, err := u.productRepository.GetProductByID(productID)
	if err != nil {
		return result, err
	}

	result = model.GetProductResponse{
		ID:             productID,
		Name:           product.Name,
		Slug:           product.Slug,
		ResellerPrice:  product.ResellerPrice,
		ConsumentPrice: product.ConsumentPrice,
		Stock:          product.Stock,
		Description:    product.Description,
		Store: model.GetStoreResponse{
			ID:       int(product.Store.ID),
			Name:     product.Store.Name,
			PhotoURL: product.Store.PhotoURL,
		},
		Category: model.GetCategoryResponse{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		},
	}

	for _, photo := range product.Photos {
		p := model.ProductPhotosResponse{
			ID:        photo.ID,
			ProductID: photo.ProductID,
			Url:       photo.Url,
		}

		result.Photos = append(result.Photos, p)
	}

	return result, nil
}

func (u *productUsecase) GetProducts(opts model.GetProductOptions) ([]model.GetProductResponse, error) {
	var result []model.GetProductResponse

	products, err := u.productRepository.GetProducts(opts)
	if err != nil {
		return result, err
	}

	for _, product := range products {
		productResponse := model.GetProductResponse{
			ID:             int(product.ID),
			Name:           product.Name,
			Slug:           product.Slug,
			ResellerPrice:  product.ResellerPrice,
			ConsumentPrice: product.ConsumentPrice,
			Stock:          product.Stock,
			Description:    product.Description,
			Store: model.GetStoreResponse{
				ID:       int(product.Store.ID),
				Name:     product.Store.Name,
				PhotoURL: product.Store.PhotoURL,
			},
			Category: model.GetCategoryResponse{
				ID:   product.Category.ID,
				Name: product.Category.Name,
			},
		}

		for _, photo := range product.Photos {
			p := model.ProductPhotosResponse{
				ID:        photo.ID,
				ProductID: photo.ProductID,
				Url:       photo.Url,
			}

			productResponse.Photos = append(productResponse.Photos, p)
		}

		result = append(result, productResponse)
	}

	return result, nil
}

func (u *productUsecase) DeleteProductByID(userID int, productID int) error {
	store, err := u.storeRepository.GetMyStore(userID)
	if err != nil {
		return err
	}

	product, err := u.productRepository.GetProductByID(productID)
	if err != nil {
		return err
	}

	if product.StoreID != store.ID {
		return fmt.Errorf("permission denied. You can only delete products from your own store")
	}

	for _, photo := range product.Photos {
		os.Remove(fmt.Sprintf("images/product_photo/%d_%s", productID, photo.Url))
	}

	return u.productRepository.DeleteProductByID(productID)
}

func (u *productUsecase) UpdateProduct(userID int, productID int, updateProductRequest model.CreateProductRequest) error {
	store, err := u.storeRepository.GetMyStore(userID)
	if err != nil {
		return err
	}

	product, err := u.productRepository.GetProductByID(productID)
	if err != nil {
		return err
	}

	if product.StoreID != store.ID {
		return fmt.Errorf("permission denied. You can only update products from your own store")
	}

	productSlug := strings.ToLower(strings.ReplaceAll(updateProductRequest.Name, " ", "-"))

	updateProduct := domain.Product{
		Model: gorm.Model{
			ID: uint(productID),
		},
		Name:           updateProductRequest.Name,
		Slug:           productSlug,
		ResellerPrice:  updateProductRequest.ResellerPrice,
		ConsumentPrice: updateProductRequest.ConsumentPrice,
		Stock:          updateProductRequest.Stock,
		Description:    updateProductRequest.Description,
		StoreID:        updateProductRequest.StoreID,
		CategoryID:     updateProductRequest.CategoryID,
	}

	for _, productPhoto := range updateProductRequest.Photos {
		updateProduct.Photos = append(updateProduct.Photos, domain.ProductPhotos{
			Url: productPhoto,
		})
	}

	return u.productRepository.UpdateProduct(updateProduct)
}
