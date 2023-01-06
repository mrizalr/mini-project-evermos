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

type UserLoginRequest struct {
	PhoneNumber string `json:"no_telp"`
	Password    string `json:"password"`
}

type UserLoginResponse struct {
	Name        string   `json:"nama"`
	PhoneNumber string   `json:"no_telp"`
	BirthDate   string   `json:"tanggal_Lahir"`
	Bio         string   `json:"tentang"`
	Job         string   `json:"pekerjaan"`
	Email       string   `json:"email"`
	ProvinceID  Province `json:"id_provinsi"`
	CityID      City     `json:"id_kota"`
	Token       string   `json:"token"`
}
