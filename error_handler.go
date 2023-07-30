package apiutils

import (
	"log"
	"net/http"
)

func HandleHttpErrors(w http.ResponseWriter, err error) {
	var code int
	var msg string
	switch err.(type) {
	case ErrBadRequest:
		code = http.StatusBadRequest
		msg = err.Error()
	case ErrNotFound:
		code = http.StatusNotFound
		msg = err.Error()
	case ErrConflict:
		code = http.StatusConflict
		msg = err.Error()
	case ErrForbidden:
		code = http.StatusForbidden
		msg = err.Error()
	case ErrUnauthorized:
		code = http.StatusUnauthorized
		msg = err.Error()
	default:
		log.Println(err)
		code = http.StatusInternalServerError
		msg = err.Error()
	}

	resp := NewRes(code, msg, nil)
	resp.SendRes(w)
}
