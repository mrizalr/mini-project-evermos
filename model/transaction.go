package model

type CreateTransactionRequest struct {
	PaymentMethod     string `json:"method_bayar"`
	AddressID         uint   `json:"alamat_kirim"`
	TransactionDetail []struct {
		ProductID uint `json:"product_id"`
		Quantity  uint `json:"kuantitas"`
	} `json:"detail_trx"`
}

type GetTransactionDetailResponse struct {
	Product    GetProductTrxResponse `json:"product"`
	Store      GetStoreResponse      `json:"toko"`
	Quantity   uint                  `json:"kuantitas"`
	TotalPrice float64               `json:"harga_total"`
}

type GetTransactionResponse struct {
	ID                uint                           `json:"id"`
	TotalPrice        float64                        `json:"total_harga"`
	Invoice           string                         `json:"kode_invoice"`
	PaymentMethod     string                         `json:"metode_bayar"`
	Address           GetAddressResponse             `json:"alamat_kirim"`
	TransactionDetail []GetTransactionDetailResponse `json:"detail_trx"`
}
