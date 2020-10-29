package router

import (
	"github.com/zerodays/sistem-users/internal/handle"
	"net/http"
)

// apiRoutes returns routes. It is in function instead of a variable to allow
// lazy loading.
func apiRoutes() []Route {
	return []Route{
		{
			Name: "authorize",
			Path: "/authorize",
			POST: http.HandlerFunc(handle.AuthorizeHandle),
		},
		{
			Name:           "change_password",
			Path:           "/change_password",
			AuthorizedOnly: true,
			PUT:            http.HandlerFunc(handle.PasswordChangeHandle),
		},

		{
			Name:           "users",
			Path:           "/users",
			AuthorizedOnly: true,
			GET:            http.HandlerFunc(handle.ListUsersHandle),
			POST:           http.HandlerFunc(handle.NewUserHandle),
		},

		{
			Name:           "user",
			Path:           "/users/{uid}",
			AuthorizedOnly: true,
			GET:            http.HandlerFunc(handle.UserHandle),
			PUT:            http.HandlerFunc(handle.EditUserHandle),
			DELETE:         http.HandlerFunc(handle.DeleteUserHandle),
		},

		{
			Name:              "signing_key",
			Path:              "/signing_key",
			CustomContentType: true,
			GET:               http.HandlerFunc(handle.SigningKeyHandle),
		},
	}
}
