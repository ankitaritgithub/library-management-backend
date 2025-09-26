package utils

import (
	"go-auth/database"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"  //for hashing and comparing passwords.
)

func ParseToken(tokenString string) (claims *database.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &database.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*database.Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
