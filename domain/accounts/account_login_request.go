package accounts

// LoginRequest is the struct to serve as interface to the login request
type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
