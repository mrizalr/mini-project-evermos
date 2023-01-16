package domain

import (
	"github.com/mrizalr/mini-project-evermos/model"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TotalPrice        float64             `json:"total_harga"`
	Invoice           string              `json:"kode_invoice" gorm:"type:varchar(30)"`
	PaymentMethod     string              `json:"metode_bayar" gorm:"type:varchar(20)"`
	AddressID         uint                `json:"id_alamat_kirim"`
	Address           Address             `json:"alamat_kirim"`
	UserID            uint                `json:"user_id"`
	TransactionDetail []TransactionDetail `json:"detail_trx"`
}

type TransactionDetail struct {
	gorm.Model
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	Product       Product `json:"product"`
	Quantity      uint    `json:"kuantitas" gorm:"int unsigned"`
	TotalPrice    float64 `json:"harga_total"`
}

type TransactionRepository interface {
	CreateNewTransaction(Transaction) (int, error)
	GetTransactionByID(int) (Transaction, error)
	GetTransactions(int) ([]Transaction, error)
}

type TransactionUsecase interface {
	CreateNewTransaction(int, model.CreateTransactionRequest) (int, error)
	GetTransactionByID(int, int) (model.GetTransactionResponse, error)
	GetTransactions(int) ([]model.GetTransactionResponse, error)
}
