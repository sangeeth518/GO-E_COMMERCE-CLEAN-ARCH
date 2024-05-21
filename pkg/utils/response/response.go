package response

type Response struct {
	StatusCode int         `json:"statuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

func ClientResponse(statuscode int, message string, data interface{}, error interface{}) Response {
	return Response{
		StatusCode: statuscode,
		Message:    message,
		Data:       data,
		Error:      error,
	}
}
