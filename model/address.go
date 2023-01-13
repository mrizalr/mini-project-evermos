package model

type GetAddressResponse struct {
	ID           uint   `json:"id"`
	Title        string `json:"judul_alamat"`
	ReceiverName string `json:"nama_penerima"`
	PhoneNumber  string `json:"no_telp"`
	Detail       string `json:"detail_alamat"`
}

type UpdateAddressRequest struct {
	ReceiverName string `json:"nama_penerima"`
	PhoneNumber  string `json:"no_telp"`
	Detail       string `json:"detail_alamat"`
}
