package user

import "github.com/zerodays/sistem-users/internal/database"

// Permissions returns slice of permissions for user.
func (u *User) Permissions() ([]string, error) {
	permissions := make([]string, 0)

	query := `SELECT permission FROM user_permissions WHERE user_id=$1`
	err := database.DB.Select(&permissions, query, u.UID)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
