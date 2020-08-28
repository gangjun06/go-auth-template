package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID int
	jwt.StandardClaims
}
