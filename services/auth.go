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
		return fmt.Errorf("token has expired")
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

func (a Auth) ParseToken(signedToken string) (*UserClaim, error) {
	t, err := jwt.ParseWithClaims(signedToken, &UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodES512.Alg() {
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
