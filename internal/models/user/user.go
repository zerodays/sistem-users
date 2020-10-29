package user

import (
	"errors"
	"github.com/google/uuid"
	"github.com/zerodays/sistem-auth/permission"
	"github.com/zerodays/sistem-users/internal/database"
	"github.com/zerodays/sistem-users/internal/logger"
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

// ForUID queries user for specified uid.
func ForUID(uid string) (*User, error) {
	u := &User{}

	query := `SELECT * FROM users WHERE uid=$1`
	err := database.DB.Get(u, query, uid)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// All returns a slice of all users.
func All() ([]*User, error) {
	var users []*User

	query := `SELECT * FROM users ORDER BY name`
	err := database.DB.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Update updates user name and their permissions.
func (u *User) Update(name string, permissions []permission.Type) error {
	// Begin transaction.
	tx, err := database.DB.Beginx()
	if err != nil {
		return err
	}

	// Set new name.
	update := `UPDATE users SET name=$2 WHERE uid=$1 RETURNING *`
	err = tx.Get(u, update, u.UID, name)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			logger.Log.Warn().Err(err).Send()
		}

		return err
	}

	// Set permissions.
	err = u.setPermissions(permissions, tx)
	if err != nil {
		return err
	}

	// Commit transaction.
	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			logger.Log.Warn().Err(err).Send()
		}

		return err
	}

	return nil
}

// Delete deletes user with specified uid.
func Delete(uid string) error {
	del := `DELETE FROM users WHERE uid=$1`
	_, err := database.DB.Exec(del, uid)
	return err
}
