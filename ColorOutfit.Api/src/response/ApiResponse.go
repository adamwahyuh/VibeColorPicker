package response

type ApiResponse[T any] struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       *T     `json:"data"`
}

func Success[T any](statusCode int, message string, data *T) ApiResponse[T] {
	return ApiResponse[T]{
		Success:    true,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

func Error[T any](statusCode int, message string) ApiResponse[T] {
	return ApiResponse[T]{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
		Data:       nil,
	}
}
