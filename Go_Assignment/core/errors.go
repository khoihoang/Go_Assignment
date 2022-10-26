package core

type ApiError struct {
	Code    int
	Message string
}

func InitApiError(code int, message string) ApiError {
	return ApiError{
		code,
		message,
	}
}
