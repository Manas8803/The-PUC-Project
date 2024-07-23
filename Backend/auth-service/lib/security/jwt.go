package security

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var Key = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT() (tokenStr string, err error) {

	claims := jwt.MapClaims{
		"exp":        time.Now().Add(720 * time.Hour).Unix(),
		"authorized": true,
		"user":       tokenStr,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err = token.SignedString(Key)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return Key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func GetExpirationTimeFromToken(tokenStr string) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return
	}

	expirationTime := time.Unix(int64(exp), 0)
	fmt.Println(expirationTime)
}
