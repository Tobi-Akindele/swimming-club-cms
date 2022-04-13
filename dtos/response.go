package dtos

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Count   int    `json:"count"`
}
