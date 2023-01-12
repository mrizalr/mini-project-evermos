package usecase

import (
	"strings"

	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/model"
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

func (u *productUsecase) CreateNewProduct(storeID int, createProductRequest model.CreateProductRequest) (int, error) {
	productSlug := strings.ToLower(strings.ReplaceAll(createProductRequest.Name, " ", "-"))

	product := domain.Product{
		Name:           createProductRequest.Name,
		Slug:           productSlug,
		ResellerPrice:  createProductRequest.ResellerPrice,
		ConsumentPrice: createProductRequest.ConsumentPrice,
		Stock:          createProductRequest.Stock,
		Description:    createProductRequest.Description,
		StoreID:        uint(storeID),
		CategoryID:     createProductRequest.CategoryID,
	}

	productPhotos := []domain.ProductPhotos{}
	for _, productPhoto := range createProductRequest.Photos {
		productPhotos = append(productPhotos, domain.ProductPhotos{
			Url: productPhoto,
		})
	}

	return u.productRepository.CreateProduct(product, productPhotos)
}

func (u *productUsecase) GetProductByID(productID int) (model.GetProductResponse, error) {
	var result model.GetProductResponse

	product, productPhotos, err := u.productRepository.GetProductByID(productID)
	if err != nil {
		return result, err
	}

	store, err := u.storeRepository.GetStoreByID(int(product.StoreID))
	if err != nil {
		return result, err
	}

	category, err := u.categoryRepository.GetCategoryByID(int(product.CategoryID))
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
			ID:       int(store.ID),
			Name:     store.Name,
			PhotoURL: store.PhotoURL,
		},
		Category: model.GetCategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		},
	}

	for _, photo := range productPhotos {
		result.Photos = append(result.Photos, photo)
	}

	return result, nil
}

func (u *productUsecase) GetProducts() ([]model.GetProductResponse, error) {
	var result []model.GetProductResponse
	productPhotos := make(map[uint][]domain.ProductPhotos)

	products, photos, err := u.productRepository.GetProducts()
	if err != nil {
		return result, err
	}

	for _, photo := range photos {
		productPhotos[photo.ProductID] = append(productPhotos[photo.ProductID], photo)
	}

	for _, product := range products {
		store, err := u.storeRepository.GetStoreByID(int(product.StoreID))
		if err != nil {
			return result, err
		}

		category, err := u.categoryRepository.GetCategoryByID(int(product.CategoryID))
		if err != nil {
			return result, err
		}

		photoResponse := []struct {
			ID        uint   `json:"id"`
			ProductID uint   `json:"product_id"`
			Url       string `json:"url"`
		}{}

		for _, productPhoto := range productPhotos[product.ID] {
			photoResponse = append(photoResponse, struct {
				ID        uint   `json:"id"`
				ProductID uint   `json:"product_id"`
				Url       string `json:"url"`
			}{
				ID:        productPhoto.ID,
				ProductID: productPhoto.ProductID,
				Url:       productPhoto.Url,
			})
		}

		productResponse := model.GetProductResponse{
			ID:             int(product.ID),
			Name:           product.Name,
			Slug:           product.Slug,
			ResellerPrice:  product.ResellerPrice,
			ConsumentPrice: product.ConsumentPrice,
			Stock:          product.Stock,
			Description:    product.Description,
			Store: model.GetStoreResponse{
				ID:       int(store.ID),
				Name:     store.Name,
				PhotoURL: store.PhotoURL,
			},
			Category: model.GetCategoryResponse{
				ID:   category.ID,
				Name: category.Name,
			},
			Photos: photoResponse,
		}

		result = append(result, productResponse)
	}

	return result, nil
}
