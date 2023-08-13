package apiutils

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"time"

	"github.com/google/uuid"
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
		Body:    body,
	}
}

func (r *Response) SendRes(w http.ResponseWriter, contentType ...string) {
	// setting content type
	contentTypeValue := "application/json"
	if len(contentType) > 0 {
		if _, _, err := mime.ParseMediaType(contentType[0]); err == nil {
			contentTypeValue = contentType[0]
		}
	}
	w.Header().Set("Content-Type", contentTypeValue)

	// setting request id
	reqId := fmt.Sprintf("%s-%s", time.Now().Format("20060102"), uuid.New().String())
	w.Header().Set("X-Request-Id", reqId)

	// setting response code
	w.WriteHeader(r.Code)

	// writing response body
	if err := json.NewEncoder(w).Encode(r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"code": 500, "message": "Internal Server Error", "body": null}`))
		return
	}
}

func (r *Response) AddHeader(w http.ResponseWriter, headers map[string]string) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
}
