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
