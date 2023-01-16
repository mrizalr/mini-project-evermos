package usecase

import (
	"fmt"
	"time"

	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/model"
)

type transactionUsecase struct {
	transactionRepository domain.TransactionRepository
	productRepository     domain.ProductRepository
}

func NewTransactionUsecase(transactionRepository domain.TransactionRepository, productRepository domain.ProductRepository) domain.TransactionUsecase {
	return &transactionUsecase{transactionRepository, productRepository}
}

func (u *transactionUsecase) CreateNewTransaction(userID int, trxRequest model.CreateTransactionRequest) (int, error) {
	invoice := fmt.Sprintf("INV-%v_%v", time.Now().Unix(), userID)

	trx := domain.Transaction{
		Invoice:       invoice,
		PaymentMethod: trxRequest.PaymentMethod,
		AddressID:     trxRequest.AddressID,
		UserID:        uint(userID),
	}

	for _, detail := range trxRequest.TransactionDetail {
		product, err := u.productRepository.GetProductByID(int(detail.ProductID))
		if err != nil {
			return 0, err
		}

		trxDetail := domain.TransactionDetail{
			ProductID:  product.ID,
			Product:    product,
			Quantity:   detail.Quantity,
			TotalPrice: float64(product.ConsumentPrice) * float64(detail.Quantity),
		}

		trx.TotalPrice += trxDetail.TotalPrice
		trx.TransactionDetail = append(trx.TransactionDetail, trxDetail)
	}

	lastInsertID, err := u.transactionRepository.CreateNewTransaction(trx)
	if err != nil {
		return lastInsertID, err
	}

	return lastInsertID, nil
}

func (u *transactionUsecase) GetTransactionByID(userID int, trxID int) (model.GetTransactionResponse, error) {
	var result model.GetTransactionResponse
	trx, err := u.transactionRepository.GetTransactionByID(trxID)
	if err != nil {
		return result, err
	}

	if trx.UserID != uint(userID) {
		return result, fmt.Errorf("permission denied. You can only get your own transaction")
	}

	result = model.GetTransactionResponse{
		ID:            trx.ID,
		TotalPrice:    trx.TotalPrice,
		Invoice:       trx.Invoice,
		PaymentMethod: trx.PaymentMethod,
		Address: model.GetAddressResponse{
			ID:           trx.Address.ID,
			Title:        trx.Address.Title,
			ReceiverName: trx.Address.ReceiverName,
			PhoneNumber:  trx.Address.PhoneNumber,
			Detail:       trx.Address.Detail,
		},
	}

	for _, trxDetail := range trx.TransactionDetail {
		trxResponse := model.GetTransactionDetailResponse{
			Product: model.GetProductTrxResponse{
				ID:             int(trxDetail.Product.ID),
				Name:           trxDetail.Product.Name,
				Slug:           trxDetail.Product.Slug,
				ResellerPrice:  trxDetail.Product.ResellerPrice,
				ConsumentPrice: trxDetail.Product.ConsumentPrice,
				Description:    trxDetail.Product.Description,
				Category: model.GetCategoryResponse{
					ID:   trxDetail.Product.Category.ID,
					Name: trxDetail.Product.Category.Name,
				},
			},
			Store: model.GetStoreResponse{
				ID:       int(trxDetail.Product.Store.ID),
				Name:     trxDetail.Product.Store.Name,
				PhotoURL: trxDetail.Product.Store.PhotoURL,
			},
			Quantity:   trxDetail.Quantity,
			TotalPrice: trxDetail.TotalPrice,
		}

		for _, photo := range trxDetail.Product.Photos {
			trxResponse.Product.Photos = append(trxResponse.Product.Photos, model.ProductPhotosResponse{
				ID:        photo.ID,
				ProductID: photo.ProductID,
				Url:       photo.Url,
			})
		}

		result.TransactionDetail = append(result.TransactionDetail, trxResponse)
	}

	return result, nil

}

func (u *transactionUsecase) GetTransactions(userID int) ([]model.GetTransactionResponse, error) {
	var result []model.GetTransactionResponse
	trx, err := u.transactionRepository.GetTransactions(userID)
	if err != nil {
		return result, err
	}

	for _, t := range trx {
		response := model.GetTransactionResponse{
			ID:            t.ID,
			TotalPrice:    t.TotalPrice,
			Invoice:       t.Invoice,
			PaymentMethod: t.PaymentMethod,
			Address: model.GetAddressResponse{
				ID:           t.Address.ID,
				Title:        t.Address.Title,
				ReceiverName: t.Address.ReceiverName,
				PhoneNumber:  t.Address.PhoneNumber,
				Detail:       t.Address.Detail,
			},
		}

		for _, trxDetail := range t.TransactionDetail {
			trxResponse := model.GetTransactionDetailResponse{
				Product: model.GetProductTrxResponse{
					ID:             int(trxDetail.Product.ID),
					Name:           trxDetail.Product.Name,
					Slug:           trxDetail.Product.Slug,
					ResellerPrice:  trxDetail.Product.ResellerPrice,
					ConsumentPrice: trxDetail.Product.ConsumentPrice,
					Description:    trxDetail.Product.Description,
					Category: model.GetCategoryResponse{
						ID:   trxDetail.Product.Category.ID,
						Name: trxDetail.Product.Category.Name,
					},
				},
				Store: model.GetStoreResponse{
					ID:       int(trxDetail.Product.Store.ID),
					Name:     trxDetail.Product.Store.Name,
					PhotoURL: trxDetail.Product.Store.PhotoURL,
				},
				Quantity:   trxDetail.Quantity,
				TotalPrice: trxDetail.TotalPrice,
			}

			for _, photo := range trxDetail.Product.Photos {
				trxResponse.Product.Photos = append(trxResponse.Product.Photos, model.ProductPhotosResponse{
					ID:        photo.ID,
					ProductID: photo.ProductID,
					Url:       photo.Url,
				})
			}

			response.TransactionDetail = append(response.TransactionDetail, trxResponse)
		}

		result = append(result, response)
	}

	return result, nil

}
