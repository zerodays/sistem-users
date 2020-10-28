package user

import (
	"github.com/zerodays/sistem-auth/permission"
	"github.com/zerodays/sistem-users/internal/database"
)

// Permissions returns slice of permissions for user.
func (u *User) Permissions() ([]permission.Permission, error) {
	permissions := make([]permission.Permission, 0)

	query := `SELECT permission FROM user_permissions WHERE user_id=$1`
	err := database.DB.Select(&permissions, query, u.UID)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
