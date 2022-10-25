package structs

type LoginRequest struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

type LoginResponse struct {
	LoggedIn     bool   `json: "loggedIn"`
	ErrorMessage string `json: "errorMessage"`
}
