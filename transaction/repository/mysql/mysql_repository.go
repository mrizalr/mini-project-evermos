package mysql

import (
	"github.com/mrizalr/mini-project-evermos/domain"
	"gorm.io/gorm"
)

type mysqlTransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &mysqlTransactionRepository{db}
}

func (r *mysqlTransactionRepository) CreateNewTransaction(trx domain.Transaction) (int, error) {
	tx := r.db.Create(&trx)
	return int(trx.ID), tx.Error
}

func (r *mysqlTransactionRepository) GetTransactionByID(trxID int) (domain.Transaction, error) {
	trx := domain.Transaction{}
	tx := r.db.Where("id = ?", trxID).
		Preload("Address").
		Preload("TransactionDetail.Product").
		Preload("TransactionDetail.Product.Store").
		Preload("TransactionDetail.Product.Category").
		Preload("TransactionDetail.Product.Photos").
		First(&trx)
	return trx, tx.Error
}

func (r *mysqlTransactionRepository) GetTransactions(userID int) ([]domain.Transaction, error) {
	trx := []domain.Transaction{}
	tx := r.db.Where("user_id = ?", userID).
		Preload("Address").
		Preload("TransactionDetail.Product").
		Preload("TransactionDetail.Product.Store").
		Preload("TransactionDetail.Product.Category").
		Preload("TransactionDetail.Product.Photos").
		Find(&trx)
	return trx, tx.Error
}
