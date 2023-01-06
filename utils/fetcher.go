package utils

import (
	"io"
	"net/http"
)

type FetchOptions struct {
	Method string
	Url    string
	Body   io.Reader
}

func Fetch(opts FetchOptions) ([]byte, error) {
	req, err := http.NewRequest(opts.Method, opts.Url, opts.Body)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}
