package security

type LoginPayload struct {
	Email    string
	Password string
}

type PasswordResetTokenPayload struct {
	Email string `json:"email"`
}

type ResetPasswordPayload struct {
	Email              string `json:"email"`
	PasswordResetToken string `json:"password_reset_token"`
	Password           string `json:"password"`
}
