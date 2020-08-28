package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gangjun06/book-server/models"
)

func GetJwtToken(id int) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * 14)

	claims := &models.Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, errStringToSignedString := token.SignedString([]byte(GetConfig().JwtSecret))

	if errStringToSignedString != nil {
		fmt.Println(errStringToSignedString)
		return "", fmt.Errorf("token signed Error")
	}
	return tokenString, nil
}
