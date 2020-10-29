package handle

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zerodays/sistem-auth/middleware"
	"github.com/zerodays/sistem-auth/permission"
	"github.com/zerodays/sistem-users/internal/handle/errors"
	"github.com/zerodays/sistem-users/internal/logger"
	"github.com/zerodays/sistem-users/internal/models/user"
	"github.com/zerodays/sistem-users/internal/util"
	"net/http"
)

type newUserRequest struct {
	Email       string                  `json:"email"`
	Name        string                  `json:"name"`
	Password    string                  `json:"password"`
	Permissions []permission.Permission `json:"permissions"`
}

type userEditRequest struct {
	Name        string                  `json:"name"`
	Permissions []permission.Permission `json:"permissions"`
}

// ListUsersHandle lists all users in database. Only users with user:read permission
// can view this page.
func ListUsersHandle(w http.ResponseWriter, r *http.Request) {
	// Get authorized user and check their permissions.
	u := middleware.UserFromRequest(r)
	if !u.HasPermission(permission.UserRead) {
		errors.Response(w, errors.InsufficientPermissions)
		return
	}

	// Get users from database.
	users, err := user.All()
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	// Write response.
	res, _ := json.Marshal(users)
	_, _ = w.Write(res)
}

// NewUserHandle creates new user.
func NewUserHandle(w http.ResponseWriter, r *http.Request) {
	// Get authorized user and check their permissions.
	authUser := middleware.UserFromRequest(r)
	if !authUser.HasPermission(permission.UserWrite) {
		errors.Response(w, errors.InsufficientPermissions)
		return
	}

	// Get data from body.
	newReq := &newUserRequest{}
	if !util.ParseJSON(w, r, newReq) {
		return
	}

	// Create new user.
	u, err := user.New(newReq.Email, newReq.Name, newReq.Password)
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	// Set user permissions.
	err = u.SetPermissions(newReq.Permissions)
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	// Write response.
	res, _ := json.Marshal(u)
	_, _ = w.Write(res)
}

// UserHandle handles getting a single user.
func UserHandle(w http.ResponseWriter, r *http.Request) {
	// Get authorized user and check their permissions.
	authUser := middleware.UserFromRequest(r)
	if !authUser.HasPermission(permission.UserRead) {
		errors.Response(w, errors.InsufficientPermissions)
		return
	}

	// Get requested user.
	uid := mux.Vars(r)["uid"]
	u, err := user.ForUID(uid)
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	// Write response.
	res, _ := json.Marshal(u)
	_, _ = w.Write(res)
}

// EditUserHandle edits specified user.
func EditUserHandle(w http.ResponseWriter, r *http.Request) {
	// Get authorized user and check their permissions.
	authUser := middleware.UserFromRequest(r)
	if !authUser.HasPermission(permission.UserWrite) {
		errors.Response(w, errors.InsufficientPermissions)
		return
	}

	// Get request user.
	uid := mux.Vars(r)["uid"]
	u, err := user.ForUID(uid)
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	// Get data from body.
	editReq := &userEditRequest{}
	if !util.ParseJSON(w, r, editReq) {
		return
	}

	// Update user.
	err = u.Update(editReq.Name, editReq.Permissions)
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	res, _ := json.Marshal(u)
	_, _ = w.Write(res)
}

// DeleteUserHandle deletes specified user.
func DeleteUserHandle(w http.ResponseWriter, r *http.Request) {
	// Get authorized user and check their permissions.
	authUser := middleware.UserFromRequest(r)
	if !authUser.HasPermission(permission.UserWrite) {
		errors.Response(w, errors.InsufficientPermissions)
		return
	}

	// Delete requested user.
	uid := mux.Vars(r)["uid"]
	err := user.Delete(uid)
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
