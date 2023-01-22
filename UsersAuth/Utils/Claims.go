package utils

import "github.com/golang-jwt/jwt/v4"

type customClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
