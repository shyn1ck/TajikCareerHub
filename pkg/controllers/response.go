package controllers

type DefaultResponse struct {
	Message string `json:"message"`
}

func NewDefaultResponse(message string) DefaultResponse {
	return DefaultResponse{
		Message: message,
	}
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}

type PasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
