package errorutils

import (
	"net/http"

	"github.com/tuan882612/apiutils"
)

func HandleHttpErrors(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *ErrBadRequest:
		body := apiutils.NewRes(http.StatusBadRequest, err.Error(), nil)
		apiutils.WrapRes(w, body)
		return
	case *ErrNotFound:
		body := apiutils.NewRes(http.StatusNotFound, err.Error(), nil)
		apiutils.WrapRes(w, body)
		return
	case *ErrConflict:
		body := apiutils.NewRes(http.StatusConflict, err.Error(), nil)
		apiutils.WrapRes(w, body)
		return
	case *ErrForbidden:
		body := apiutils.NewRes(http.StatusForbidden, err.Error(), nil)
		apiutils.WrapRes(w, body)
		return
	case *ErrUnauthorized:
		body := apiutils.NewRes(http.StatusUnauthorized, err.Error(), nil)
		apiutils.WrapRes(w, body)
		return
	default:
		body := apiutils.NewRes(http.StatusInternalServerError, err.Error(), nil)
		apiutils.WrapRes(w, body)
		return
	}
}
