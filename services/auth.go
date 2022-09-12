package services

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type IAuth interface {
	CreateToken(c *UserClaim) (string, error)
	ParseToken(signedToken string) (*UserClaim, error)
}

type Auth struct {
	AuthKey []byte
}

type UserClaim struct {
	jwt.StandardClaims
	Email string
}

func (a Auth) CreateToken(c *UserClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	jt, err := token.SignedString(a.AuthKey)
	if err != nil {
		return "", fmt.Errorf("error while signing")
	}
	return jt, nil
}

func (a Auth) ParseToken(signedToken string) (*UserClaim, error) {
	t, err := jwt.ParseWithClaims(signedToken, &UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("invalid signature algorithm")
		}
		return a.AuthKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error in ParseToken while parsing token: %w", err)
	}
	if !t.Valid {
		return nil, fmt.Errorf("error in ParseToken, token is not valid")
	}
	return t.Claims.(*UserClaim), nil
}
