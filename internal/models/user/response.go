package user

import (
	"encoding/json"
	"github.com/zerodays/sistem-auth/permission"
	"github.com/zerodays/sistem-users/internal/logger"
)

type Response struct {
	UID string `json:"id"`

	Email string `json:"email"`
	Name  string `json:"name"`

	Permissions []permission.Permission `json:"permissions"`
}

func (u *User) Response() Response {
	permissions, err := u.Permissions()
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		permissions = []permission.Permission{}
	}

	return Response{
		UID:         u.UID,
		Email:       u.Email,
		Name:        u.Name,
		Permissions: permissions,
	}
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Response())
}
