package common

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}