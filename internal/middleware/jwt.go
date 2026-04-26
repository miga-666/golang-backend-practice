package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"ginBackend/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Validates a Bearer JWT in the Authorization header.
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			abortUnauthorized(ctx, "缺少或格式錯誤的 JWT Token")
			return
		}

		// Strip the "Bearer " prefix to obtain the raw token string.
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		secret := os.Getenv("JWT_SECRET")
		claims := &model.JWTClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Reject tokens signed with any algorithm other than HMAC
			// to prevent the "alg:none" and RSA confusion attacks.
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			abortUnauthorized(ctx, "無效或已過期的 JWT Token")
			return
		}

		// Forward parsed claims to downstream handlers.
		ctx.Set("email", claims.Email)
		ctx.Set("updated", claims.Updated)
		ctx.Next()
	}
}

// writes a 401 JSON response
func abortUnauthorized(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.APIResponse{
		Success:   false,
		Message:   message,
		Data:      nil,
		Error:     nil,
		Code:      http.StatusUnauthorized,
		Timestamp: time.Now().UTC(),
	})
}
