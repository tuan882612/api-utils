package apiutils

import (
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func NewRes(code int, message string, body interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Body:    &body,
	}
}

func WrapRes(w http.ResponseWriter, value *Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(value.Code)
	Jsonify(w, value)
}