package model

type UserRegisterRequest struct {
	Name        string `json:"nama"`
	Password    string `json:"kata_sandi"`
	PhoneNumber string `json:"no_telp"`
	Birthdate   string `json:"tanggal_lahir"`
	Job         string `json:"pekerjaan"`
	Email       string `json:"email"`
	ProvinceID  string `json:"id_provinsi"`
	CityID      string `json:"id_kota"`
}
