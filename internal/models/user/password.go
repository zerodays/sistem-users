package user

import (
	"github.com/zerodays/sistem-users/internal/database"
	"golang.org/x/crypto/bcrypt"
)

// SetPassword sets password for specified user.
func SetPassword(uid, password string) error {
	// Hash user password.
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return err
	}

	// Save new password to database.
	update := `UPDATE users SET password=$2 WHERE uid=$1`
	_, err = database.DB.Exec(update, uid, string(passwordHash))
	return err
}

// PasswordMatch checks if user has specified password.
func (u *User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
