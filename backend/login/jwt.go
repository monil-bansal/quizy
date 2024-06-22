package login

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey = []byte("secret-keyWHATsoEVER")
)

/*
	This class is copy pasted from :
		https://github.com/cheildo/jwt-auth-golang/blob/master/login/svc.go
		https://medium.com/@cheickzida/golang-implementing-jwt-token-authentication-bba9bfd84d60
*/

func CreateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": userId,
			// expiring in 24 hours
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		return "", err
	}
	if time.Now().After(exp.Time) {
		return "", fmt.Errorf("expired token")
	}
	iss, err := token.Claims.GetIssuer()

	if err != nil {
		return "", err
	}
	return iss, nil
}
