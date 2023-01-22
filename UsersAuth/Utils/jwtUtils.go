package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwtToken(emailId string) (string, error) {
	expirationTime := time.Now().Add(time.Minute * 1)
	newClaim := customClaims{
		Email:            emailId,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: &jwt.NumericDate{Time: expirationTime}},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaim)
	fmt.Println(os.Getenv("SECRET_KEY"))
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func ValidateJwtToken(token string) (*customClaims, bool, error) {
	tkn, err := jwt.ParseWithClaims(token, &customClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {

		return &customClaims{}, false, err

	}
	return tkn.Claims.(*customClaims), tkn.Valid, nil
}

func RefreshJwtToken(token string) (string, error) {
	claims, valid, err := ValidateJwtToken(token)

	if !valid {
		if verror, ok := err.(*jwt.ValidationError); ok && verror.Errors != jwt.ValidationErrorExpired {
			return "", errors.New("INVALID TOKEN")
		}
	}
	if valid {
		return token, nil
	}
	expirationTime := time.Now().Add(time.Minute * 1)
	claims.ExpiresAt = &jwt.NumericDate{Time: expirationTime}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return newToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

}
