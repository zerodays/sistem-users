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
			Name:              "signing_key",
			Path:              "/signing_key",
			CustomContentType: true,
			GET:               http.HandlerFunc(handle.SigningKeyHandle),
		},
	}
}
