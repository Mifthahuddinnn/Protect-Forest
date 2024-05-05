package constant

import "errors"

var (
	ErrorInvalidInput = errors.New("invalid Add Data in Database")
	ErrorEmptyInput   = errors.New("name, email and password are required")
	ErrorEmailInvalid = errors.New("email is invalid")
	ErrorPassword     = errors.New("password must be at least 6 characters")
	ErrorEmailExists  = errors.New("email already exists")
	ErrorToken        = errors.New("failed to create token")
	ErrorNotFound     = errors.New("data not found")
	ErrorInternal     = errors.New("internal server error")
	ErrorUnauthorized = errors.New("unauthorized")
)
