package utils

import (
	"encoding/json"
)

func GetRegionData(url string, result interface{}) error {
	data, err := Fetch(FetchOptions{
		Method: "GET",
		Url:    url,
		Body:   nil,
	})

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return err
	}

	return nil
}
