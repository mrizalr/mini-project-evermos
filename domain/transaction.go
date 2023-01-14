package domain

import (
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TotalPrice        float64 `json:"total_harga"`
	Invoice           string  `json:"kode_invoice"`
	PaymentMethod     string  `json:"metode_bayar"`
	AddressID         uint    `json:"id_alamat_kirim"`
	Address           Address
	TransactionDetail []TransactionDetail `json:"detail_trx"`
}

type TransactionDetail struct {
	gorm.Model
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	Product       Product `json:"product"`
	Quantity      uint    `json:"kuantitas"`
	TotalPrice    float64 `json:"harga_total"`
}

type TransactionRepository interface {
	CreateNewTransaction(Transaction) (int, error)
}

type TransactionUsecase interface {
	CreateNewTransaction(int, model.CreateTransactionRequest) (int, error)
}
