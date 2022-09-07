package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	AuthKey []byte
}

type UserClaim struct {
	jwt.StandardClaims
	SessionId int64
}

func (u *UserClaim) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token has expired")
	}

	if u.SessionId == 0 {
		return fmt.Errorf("invalid session Id")
	}

	return nil
}

func (a Auth) CreateToken(c *UserClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	jt, err := token.SignedString(a.AuthKey)
	if err != nil {
		return "", fmt.Errorf("error while signing")
	}
	return jt, nil
}
