package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/zerodays/sistem-users/internal/config"
	"time"
)

type TokenClaims struct {
	jwt.StandardClaims

	Permissions []string `json:"permissions"`
}

// CreateAccessToken creates new access token and signs it.
func (u *User) CreateAccessToken() (string, error) {
	permission, err := u.Permissions()
	if err != nil {
		return "", err
	}

	now := time.Now()
	expiresAt := now.Add(time.Duration(config.Login.TokenTtl) * time.Second)

	claims := TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
			Subject:   u.UID,
		},

		Permissions: permission,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(config.Login.SigningPrivateKey)

	return ss, err
}
