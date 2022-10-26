package core

import "fmt"

type ApiResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateSuccessResponse(data interface{}) ApiResponse {
	return ApiResponse{
		Success: false,
		Code:    200,
		Message: "",
		Data:    data,
	}
}

func CreateApiErrorResponse(apiError ApiError, data interface{}) ApiResponse {
	return ApiResponse{
		Success: false,
		Code:    apiError.Code,
		Message: apiError.Message,
		Data:    data,
	}
}

func CreateUserApiErrorResponse(apiError ApiError, additionalMessage string, data interface{}) ApiResponse {
	return ApiResponse{
		Success: false,
		Code:    apiError.Code,
		Message: fmt.Sprintf("%v|%v", apiError.Message, additionalMessage),
		Data:    data,
	}
}
