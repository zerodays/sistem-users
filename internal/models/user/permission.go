package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/zerodays/sistem-auth/permission"
	"github.com/zerodays/sistem-users/internal/database"
	"github.com/zerodays/sistem-users/internal/logger"
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

// setPermissions sets user permission with transaction.
func (u *User) setPermissions(permissions []permission.Permission, tx *sqlx.Tx) error {
	// Remove current permissions.
	del := `DELETE FROM user_permissions WHERE user_id=$1`
	_, err := tx.Exec(del, u.UID)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			logger.Log.Warn().Err(err).Send()
		}

		return err
	}

	// Add permissions.
	for _, perm := range permissions {
		insert := `INSERT INTO user_permissions (user_id, permission) VALUES ($1, $2)`
		_, err := tx.Exec(insert, u.UID, perm)
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				logger.Log.Warn().Err(err).Send()
			}

			return err
		}
	}

	return nil
}

func (u *User) SetPermissions(permissions []permission.Permission) error {
	// Create new transaction.
	tx, err := database.DB.Beginx()
	if err != nil {
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
