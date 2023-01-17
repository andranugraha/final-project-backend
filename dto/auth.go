package dto

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type SignInResponse struct {
	AccessToken string `json:"accessToken"`
}
