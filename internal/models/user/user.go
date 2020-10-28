package user

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/zerodays/sistem-users/internal/database"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const bcryptCost = 12

var ErrInvalidCredentials = errors.New("invalid credentials")

type User struct {
	UID string `db:"uid"`

	Email string `db:"email"`
	Name  string `db:"name"`

	Password string `db:"password"`

	DateCreated time.Time `db:"date_created"`
}

// New creates new User instance with given email, name and password
// and saves the user to database. Initial user has no permissions
// and no authentication tokens.
func New(email, name, password string) (*User, error) {
	// Generate UID for user.
	uid := uuid.New().String()

	// Hash user password.
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return nil, err
	}

	// Insert user to database.
	u := &User{}
	insert := `INSERT INTO users (uid, email, name, password) VALUES ($1, $2, $3, $4) RETURNING *`
	err = database.DB.Get(u, insert, uid, email, name, string(passwordHash))

	// Handle errors.
	if err != nil {
		return nil, err
	}

	return u, nil
}

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
