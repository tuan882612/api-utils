package apiutils

import "net/http"

func HandleHttpErrors(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *ErrBadRequest:
		body := NewRes(http.StatusBadRequest, err.Error(), nil)
		WrapRes(w, body)
		return
	case *ErrNotFound:
		body := NewRes(http.StatusNotFound, err.Error(), nil)
		WrapRes(w, body)
		return
	case *ErrConflict:
		body := NewRes(http.StatusConflict, err.Error(), nil)
		WrapRes(w, body)
		return
	case *ErrForbidden:
		body := NewRes(http.StatusForbidden, err.Error(), nil)
		WrapRes(w, body)
		return
	case *ErrUnauthorized:
		body := NewRes(http.StatusUnauthorized, err.Error(), nil)
		WrapRes(w, body)
		return
	default:
		body := NewRes(http.StatusInternalServerError, err.Error(), nil)
		WrapRes(w, body)
		return
	}
}
