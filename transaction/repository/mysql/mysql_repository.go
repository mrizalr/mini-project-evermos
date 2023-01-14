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
