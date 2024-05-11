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
	ErrorOpenFile     = errors.New("failed to open file")
	ErrorUploadFile   = errors.New("failed to upload file to Cloudinary")
	ErrorInitCloud    = errors.New("failed to initialize Cloudinary service")
	ErrorGetFile      = errors.New("failed to get file from form")
	ErrorFailidCreate = errors.New("failed to create report")
	ErrorInvalidToken = errors.New("invalid user token")
)
