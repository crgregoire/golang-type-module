package types

//
// Login is used for logging in via Cognito
//
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
