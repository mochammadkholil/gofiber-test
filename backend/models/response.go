package models

type ApiResponse struct {
	Success bool
	Message string
	Data    interface{}
}
