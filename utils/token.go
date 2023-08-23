package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenService struct{}

func (s JWTTokenService) GenerateToken(email string) (string, error) {
	Scrt := os.Getenv("SECTRET_KEY")
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Scrt))
}

func (s JWTTokenService) ValidateToken(tokenstring string) (string, error) {
	Scrt := os.Getenv("SECTRET_KEY")
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Scrt), nil
	})

	if err != nil {
		return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		return email, nil
	}
	return "", fmt.Errorf("Error extracting email from token")
}
