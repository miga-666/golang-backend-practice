package service

import (
	"os"
	"time"

	"ginBackend/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT creates a signed HS256 JWT using model.JWTClaims.
// The secret is read from the JWT_SECRET environment variable at call time.
// No expiry is configured — tokens are invalidated by password changes via the Updated field.
func generateJWT(email string, updated time.Time) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := model.JWTClaims{
		Email:   email,
		Updated: updated,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
