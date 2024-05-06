package utils

import (
	"forest/constant"
	"net/http"
)

func ConvertResponseCode(err error) int {
	switch err {
	case constant.ErrorEmptyInput:
		return http.StatusBadRequest
	case constant.ErrorEmailInvalid:
		return http.StatusBadRequest
	case constant.ErrorPassword:
		return http.StatusBadRequest
	case constant.ErrorEmailExists:
		return http.StatusBadRequest
	case constant.ErrorToken:
		return http.StatusInternalServerError
	case constant.ErrorNotFound:
		return http.StatusNotFound
	case constant.ErrorUnauthorized:
		return http.StatusUnauthorized
	case constant.ErrorInvalidInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
