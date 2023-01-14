package model

type CreateTransactionRequest struct {
	PaymentMethod     string `json:"method_bayar"`
	AddressID         uint   `json:"alamat_kirim"`
	TransactionDetail []struct {
		ProductID uint `json:"product_id"`
		Quantity  uint `json:"kuantitas"`
	} `json:"detail_trx"`
}
