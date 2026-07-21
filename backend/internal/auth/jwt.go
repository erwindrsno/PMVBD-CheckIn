package auth

import (
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Helper to grab secret safely (assuming .env is loaded in main)
func getSecret() []byte {
	// Optional fallback if you want safety during development
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		slog.Warn("JWT_SECRET is empty, using fallback")
		return []byte("fallback-secret-key")
	}
	return []byte(secret)
}

func GenerateToken() (string, error) {
	if err := godotenv.Load(); err != nil {
		slog.Error("error loading .env")
	}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getSecret())
}

func ValidateToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return getSecret(), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
