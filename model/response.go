package model

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors"`
	Data    interface{} `json:"data"`
}
