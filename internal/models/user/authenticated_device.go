package user

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/zerodays/sistem-users/internal/database"
	"time"
)

type AuthenticatedDevice struct {
	UserUID string `db:"user_id"`

	Token string `db:"token"`

	DateCreated time.Time `db:"date_created"`
	LastUsed    time.Time `db:"last_used"`
}

// NewAuthorizedDevice creates new authorized device for user and saves id to DB.
func (u *User) NewAuthorizedDevice() (*AuthenticatedDevice, error) {
	// Create token.
	tokenBytes := make([]byte, 128)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return nil, err
	}

	token := base64.URLEncoding.EncodeToString(tokenBytes)

	// Insert into database.
	dev := &AuthenticatedDevice{}
	insert := `INSERT INTO authenticated_devices (user_id, token, last_used) VALUES ($1, $2, $3) RETURNING *`
	err = database.DB.Get(dev, insert, u.UID, token, time.Now())
	if err != nil {
		return nil, err
	}

	return dev, nil
}
