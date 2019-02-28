package handlers

import (
	"log"
	"net/http"

	"github.com/hitesh-goel/loomx/internal/platform/db"
	"github.com/hitesh-goel/loomx/internal/platform/web"
)

// API returns a handler for a set of routes.
func API(l *log.Logger, masterDB *db.DB) http.Handler {
	// TODO: User Middlewares for tracing
	app := web.New(l)

	h := Health{}
	u := User{
		MasterDB: masterDB,
	}
	app.Handle("GET", "/v1/health", h.Check)
	app.Handle("POST", "/v1/user", u.Create)
	app.Handle("GET", "/v1/user", u.Retrieve)
	app.Handle("GET", "/v1/user/:id", u.Retrieve)

	return app
}
