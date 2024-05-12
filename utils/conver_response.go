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
	case constant.ErrorOpenFile:
		return http.StatusBadRequest
	case constant.ErrorUploadFile:
		return http.StatusInternalServerError
	case constant.ErrorInitCloud:
		return http.StatusInternalServerError
	case constant.ErrorTitleEmpty:
		return http.StatusBadRequest
	case constant.ErrorContentEmpty:
		return http.StatusBadRequest
	case constant.ErrorFailedCreate:
		return http.StatusInternalServerError
	case constant.ErrorUploadFile:
		return http.StatusInternalServerError
	case constant.ErrorReportAlreadyApproved:
		return http.StatusBadRequest
	case constant.ErrorUsernameAlreadyExist:
		return http.StatusBadRequest
	case constant.ErrorAdminEmptyField:
		return http.StatusBadRequest
		case constant.ErrorReportNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
