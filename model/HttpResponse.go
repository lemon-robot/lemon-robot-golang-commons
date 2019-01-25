package model

type HttpResponse struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
}
