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
			Name:              "test",
			Path:              "/test",
			CustomContentType: true,
			GET:               http.HandlerFunc(handle.TestHandle),
		},
	}
}
