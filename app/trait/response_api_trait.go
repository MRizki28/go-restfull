package trait

type ResponseAPI struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func SuccessResponse(message string, data interface{}) ResponseAPI {
	return ResponseAPI{
		Status:  "Success",
		Message: message,
		Data:    data,
	}
}

func DataNotFoundResponse(message string) ResponseAPI {
	return ResponseAPI{
		Status:  "Not Found",
		Message: message,
	}
}

func DeleteDataSuccessResponse(message string) ResponseAPI {
	return ResponseAPI{
		Status:  "Success",
		Message: message,
	}
}

func ErrorResponse(status string, message string, err interface{}) ResponseAPI {
	return ResponseAPI{
		Status:  status,
		Message: message,
		Error:   err,
	}
}