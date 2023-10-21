package utils

type WebResponse struct {
	Code   int32       `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Code    int32       `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}
