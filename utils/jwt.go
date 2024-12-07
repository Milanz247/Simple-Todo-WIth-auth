package utils

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte("UzI1NiIsInR5cCI6IkpXVCJ9")

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

	
// VerifyJWT decodes and verifies the JWT token
func VerifyJWT(tokenStr string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Validate token
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
	func JWTMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
				return
			}
	
			// Extract the token from "Bearer <token>"
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenStr == authHeader {
				http.Error(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
				return
			}
	
			// Verify the JWT token
			claims, err := VerifyJWT(tokenStr)
			if err != nil {
				http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
				return
			}
	
			// Validate `exp` claim (if not handled in VerifyJWT)
			if exp, ok := claims["exp"].(float64); ok {
				if time.Unix(int64(exp), 0).Before(time.Now()) {
					http.Error(w, "Unauthorized: Token expired", http.StatusUnauthorized)
					return
				}
			}
	
			// Add user ID to context
			userID, ok := claims["user_id"].(float64)
			if !ok {
				http.Error(w, "Unauthorized: Invalid user ID", http.StatusUnauthorized)
				return
			}
	
			// Add the user ID to the request context
			ctx := context.WithValue(r.Context(), "user_id", uint(userID))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	