package controllers

type defaultResponse struct {
	Message string `json:"message"`
}

func newDefaultResponse(message string) defaultResponse {
	return defaultResponse{
		Message: message,
	}
}

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}
