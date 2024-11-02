package types

type RegisterRequest struct {
	FirstName string `json:"firstName" validate:"required,min=3,max=80"`
	LastName  string `json:"lastName" validate:"required,min=3,max=80"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,strongpassword"`
	Role      string `json:"role" validate:"omitempty,oneof=seller customer"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,strongpassword"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
