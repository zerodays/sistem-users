package user

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/zerodays/sistem-auth/token"
	"github.com/zerodays/sistem-users/internal/config"
	"github.com/zerodays/sistem-users/internal/database"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// AuthorizeWithPassword authorizes user with provided email and password.
// If email or password is incorrect, ErrInvalidCredentials is returned.
func AuthorizeWithPassword(email, password string) (*User, error) {
	u := &User{}

	// Get user for specified email.
	query := `SELECT * FROM users WHERE email=$1`
	err := database.DB.Get(u, query, email)

	// Handle errors.
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvalidCredentials
		} else {
			return nil, err
		}
	}

	// Check that passwords match.
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return u, nil
}

// AuthorizeWithToken authorizes user with provided token.
// If user for token does not exist, ErrInvalidCredentials is returned.
func AuthorizeWithToken(token string) (*User, error) {
	u := &User{}

	// Get user with token.
	query := `SELECT users.* FROM authenticated_devices INNER JOIN users ON authenticated_devices.user_id = users.uid
		WHERE token=$1`
	err := database.DB.Get(u, query, token)

	// Handle errors.
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvalidCredentials
		} else {
			return nil, err
		}
	}

	return u, nil
}

// CreateAccessToken creates new access token and signs it.
func (u *User) CreateAccessToken() (string, error) {
	permission, err := u.Permissions()
	if err != nil {
		return "", err
	}

	now := time.Now()
	expiresAt := now.Add(time.Duration(config.Login.TokenTtl()) * time.Second)

	claims := token.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
			Subject:   u.UID,
		},

		Permissions: permission,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := accessToken.SignedString(config.Login.SigningPrivateKey)

	return ss, err
}
