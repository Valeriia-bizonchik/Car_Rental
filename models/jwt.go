package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserID uint `json:"user_id"`
	Role   Role `json:"role"`
	jwt.StandardClaims
}
