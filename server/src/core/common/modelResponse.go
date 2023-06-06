package common

type ResponseData interface {
}

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"statusCode"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
	Error   interface{} `json:"error"`
}
