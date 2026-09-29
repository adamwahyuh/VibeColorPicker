package response

import "net/http"

func ErrorNotFound[T any]() ApiResponse[T] {
	return Error[T](http.StatusNotFound, "Not found")
}

func ErrorUnauthorized[T any]() ApiResponse[T] {
	return Error[T](http.StatusUnauthorized, "Unauthorize.")
}

func ErrorForbidden[T any]() ApiResponse[T] {
	return Error[T](http.StatusForbidden, "Access forbidden.")
}

func ErrorInternalServer[T any]() ApiResponse[T] {
	return Error[T](http.StatusInternalServerError, "Internal server error.")
}
