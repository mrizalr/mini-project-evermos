package utils

import (
	"encoding/json"
	"fmt"

	"github.com/mrizalr/mini-project-evermos/model"
)

var baseUrl string = `https://www.emsifa.com/api-wilayah-indonesia/api`

func GetProvince(ProvinceID uint) (model.Province, error) {
	url := fmt.Sprintf(`%s/province/%d.json`, baseUrl, ProvinceID)
	province := model.Province{}

	data, err := Fetch(FetchOptions{
		Method: "GET",
		Url:    url,
		Body:   nil,
	})

	if err != nil {
		return province, err
	}

	err = json.Unmarshal(data, &province)
	if err != nil {
		return province, err
	}

	return province, nil
}

func GetCity(CityID uint) (model.City, error) {
	url := fmt.Sprintf(`%s/regency/%d.json`, baseUrl, CityID)
	city := model.City{}

	data, err := Fetch(FetchOptions{
		Method: "GET",
		Url:    url,
		Body:   nil,
	})

	if err != nil {
		return city, err
	}

	err = json.Unmarshal(data, &city)
	if err != nil {
		return city, err
	}

	return city, nil
}
