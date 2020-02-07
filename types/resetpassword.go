package types

//
// ResetPassword is used for Vijnana's reset
// password route
//
type ResetPassword struct {
	Session     string `json:"session"`
	NewPassword string `json:"newPassword"`
	Email       string `json:"email"`
}
