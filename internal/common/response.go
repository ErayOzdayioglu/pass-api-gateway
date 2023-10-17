package common

type Response struct {
	Result  any           `json:"result,omitempty"`
	Message string        `json:"message,omitempty"`
	Error   errorResponse `json:"error,omitempty"`
}

type errorResponse struct {
	Message string `json:"message,omitempty"`
}

func SuccessResponse(data any) *Response {
	return &Response{
		Message: "Success",
		Result:  data,
	}
}

func ErrorResponse(message string) *Response {
	return &Response{
		Message: "Error",
		Error: errorResponse{
			Message: message,
		},
	}

}
