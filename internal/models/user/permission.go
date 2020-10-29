package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/zerodays/sistem-auth/permission"
	"github.com/zerodays/sistem-users/internal/database"
	"github.com/zerodays/sistem-users/internal/logger"
)

// Permissions returns slice of permissions for user.
func (u *User) Permissions() ([]permission.Type, error) {
	permissions := make([]permission.Type, 0)

	query := `SELECT permission FROM user_permissions WHERE user_id=$1`
	err := database.DB.Select(&permissions, query, u.UID)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

// addPermission adds permissions to user using current transaction.
// It also handles implicit permissions.
func (u *User) addPermission(perm permission.Type, tx *sqlx.Tx) error {
	// Add implicit permission.
	for _, implicitPerm := range perm.ImplicitPermissions {
		err := u.addPermission(implicitPerm, tx)
		if err != nil {
			return err
		}
	}

	// Add the permission.
	insert := `INSERT INTO user_permissions (user_id, permission) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := tx.Exec(insert, u.UID, perm)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			logger.Log.Warn().Err(err).Send()
		}

		return err
	}

	return nil
}

// setPermissions sets user permission with transaction.
func (u *User) setPermissions(permissions []permission.Type, tx *sqlx.Tx) error {
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
		err = u.addPermission(perm, tx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *User) SetPermissions(permissions []permission.Type) error {
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
