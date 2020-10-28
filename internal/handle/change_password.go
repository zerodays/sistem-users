package handle

import (
	"fmt"
	"github.com/zerodays/sistem-auth/middleware"
	"github.com/zerodays/sistem-auth/permission"
	"github.com/zerodays/sistem-users/internal/handle/errors"
	"github.com/zerodays/sistem-users/internal/logger"
	"github.com/zerodays/sistem-users/internal/models/user"
	"github.com/zerodays/sistem-users/internal/util"
	"net/http"
)

type passwordChangeRequest struct {
	UserID      string `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// PasswordChangeHandle handles changing user password. If user has user:write permission
// they can change password for any user, otherwise they can only change their password.
func PasswordChangeHandle(w http.ResponseWriter, r *http.Request) {
	// Get data from request.
	passRequest := &passwordChangeRequest{}
	if !util.ParseJSON(w, r, passRequest) {
		return
	}

	// Get authenticated user.
	authUser := middleware.UserFromRequest(r)
	fmt.Println(authUser.Permissions)
	if authUser.HasPermission(permission.UserWrite) {
		// Change password of specified user, since they have write permission.
		err := user.SetPassword(passRequest.UserID, passRequest.NewPassword)
		if err != nil {
			logger.Log.Warn().Err(err).Send()

			errors.Response(w, errors.DatabaseError)
			return
		}
	} else {
		// User can only check their own password.
		// Check if authenticated user is the user for which the request was made.
		if authUser.UID != passRequest.UserID {
			errors.Response(w, errors.InvalidCredentials)
			return
		}

		// Get user object from database.
		u, err := user.ForUID(authUser.UID)
		if err != nil {
			logger.Log.Warn().Err(err).Send()

			errors.Response(w, errors.DatabaseError)
			return
		}

		// Check if old password is correct.
		if !u.PasswordMatch(passRequest.OldPassword) {
			errors.Response(w, errors.InvalidOldPassword)
			return
		}

		// Set new password.
		err = user.SetPassword(u.UID, passRequest.NewPassword)
		if err != nil {
			logger.Log.Warn().Err(err).Send()

			errors.Response(w, errors.DatabaseError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
