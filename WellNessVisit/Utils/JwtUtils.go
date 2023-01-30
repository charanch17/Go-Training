package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwtToken(uid uint) (string, error) {
	expirationTime := time.Now().AddDate(0, 0, 7)
	claims := CustomClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
			Issuer:    "wellnesvisit.com",
		},
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return Token.SignedString([]byte(os.Getenv("SECRET_KEY")))

}

func ValidateJwtToken(tokenString string) (*CustomClaims, error) {
	Token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return &CustomClaims{}, err
	}
	return Token.Claims.(*CustomClaims), nil
}

func RefreshJwtToken(tokenString string) (string, error) {
	claims, err := ValidateJwtToken(tokenString)
	if err != nil {
		return "", err
	}
	remainingTime := time.Until(claims.ExpiresAt.Time).Hours()
	if remainingTime <= 24 {
		return GenerateJwtToken(claims.Uid)
	}
	return tokenString, nil

}
