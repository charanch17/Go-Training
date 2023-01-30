package utils

import (
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Uid uint
	jwt.RegisteredClaims
}
