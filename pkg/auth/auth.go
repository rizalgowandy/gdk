package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidType  = errors.New("invalid token type")
)

type Operator struct {
	Enable      bool
	JWTSecret   string
	JWTDuration time.Duration
}

// GenerateToken creates a new JWT token for a user
func (h *Operator) GenerateToken(user Claims) (TokenResponse, error) {
	expirationTime := time.Now().Add(h.JWTDuration)

	// Create the JWT claims
	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"email":   user.Email,
		"exp":     expirationTime.Unix(),
		"iat":     time.Now().Unix(),
	}

	// Create a token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	tokenString, err := token.SignedString([]byte(h.JWTSecret))
	if err != nil {
		return TokenResponse{}, fmt.Errorf("signing JWT token: %w", err)
	}

	return TokenResponse{
		Token:     tokenString,
		ExpiresAt: expirationTime,
	}, nil
}

func (h *Operator) GenerateTokenWithTTL(
	user Claims,
	ttl time.Duration,
) (TokenResponse, error) {
	if ttl.Seconds() <= 0 {
		return TokenResponse{}, fmt.Errorf("invalid TTL: %w", ErrInvalidType)
	}

	expirationTime := time.Now().Add(ttl)

	// Create the JWT claims
	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"email":   user.Email,
		"exp":     expirationTime.Unix(),
		"iat":     time.Now().Unix(),
	}

	// Create a token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	tokenString, err := token.SignedString([]byte(h.JWTSecret))
	if err != nil {
		return TokenResponse{}, fmt.Errorf("signing JWT token: %w", err)
	}

	return TokenResponse{
		Token:     tokenString,
		ExpiresAt: expirationTime,
	}, nil
}

// ValidateToken validates and parses the JWT token
func (h *Operator) ValidateToken(tokenString string) (Claims, error) {
	if !h.Enable {
		return Claims{
			UserID: 1,
			Email:  "admin@gmail.com",
		}, nil
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(
				"unexpected signing method: %v: %w",
				token.Header["alg"],
				ErrInvalidToken,
			)
		}
		return []byte(h.JWTSecret), nil
	})

	if err != nil {
		return Claims{}, fmt.Errorf("parsing JWT token: %w", err)
	}

	// Validate token
	if !token.Valid {
		return Claims{}, ErrInvalidToken
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Claims{}, fmt.Errorf("extracting JWT claims: %w", ErrInvalidToken)
	}

	// Check expiration
	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		return Claims{}, ErrExpiredToken
	}

	// Extract user claims
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return Claims{}, fmt.Errorf("claims user id: %w", ErrInvalidType)
	}
	email, ok := claims["email"].(string)
	if !ok {
		return Claims{}, fmt.Errorf("claims email: %w", ErrInvalidType)
	}

	return Claims{
		UserID: int(userID),
		Email:  email,
	}, nil
}

func NewOperator(enable bool, secret string, duration time.Duration) *Operator {
	return &Operator{
		Enable:      enable,
		JWTSecret:   secret,
		JWTDuration: duration,
	}
}
