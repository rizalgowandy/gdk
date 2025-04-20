package auth

import "time"

// TokenResponse represents the response after successful login
type TokenResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Claims represent the JWT token claims
type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}
