package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims is the shared JWT payload used when signing (service) and
// verifying (middleware) tokens. Centralising it here ensures both packages
// always agree on the claim structure.
type JWTClaims struct {
	Email   string    `json:"email"`
	Updated time.Time `json:"updated"`
	jwt.RegisteredClaims
}

// All responses — success and error alike — use this structure so that
type APIResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`  // Populated on success, null on error.
	Error     interface{} `json:"error"` // Populated on error, null on success.
	Code      int         `json:"code"`
	Timestamp time.Time   `json:"timestamp"`
}

type RegisterResponse struct {
	Email string `json:"email"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"` // Returned JWT token when login is successful.
}
