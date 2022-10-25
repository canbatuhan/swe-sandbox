package structs

type RegisterRequest struct {
	Email           string `json: "email"`
	Username        string `json: "username"`
	Password        string `json: "password"`
	ConfirmPassword string `json: "confirmPassword"`
}

type RegisterResponse struct {
	Registered   bool   `json: "registered"`
	ErrorMessage string `json: "errorMessage"`
}
