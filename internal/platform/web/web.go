package web

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dimfeld/httptreemux"
	uuid "github.com/satori/go.uuid"
)

// ctxKey represents the type of value for the context key.
type ctxKey int

// KeyValues is how request values or stored/retrieved.
const KeyValues ctxKey = 1

// Values represent state for each request.
type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
	Error      bool
}

// A Handler is a type that handles an http request within our own little mini
// framework.
type Handler func(ctx context.Context, log *log.Logger, w http.ResponseWriter, r *http.Request, params map[string]string) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct
type App struct {
	*httptreemux.TreeMux
	log *log.Logger
}

// New creates an App value that handle a set of routes for the application.
func New(log *log.Logger) *App {
	return &App{
		TreeMux: httptreemux.New(),
		log:     log,
	}
}

// Handle is our mechanism for mounting Handlers for a given HTTP verb and path
// pair, this makes for really easy, convenient routing.
func (a *App) Handle(verb, path string, handler Handler) {

	h := func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		ctx := context.TODO()

		u, err := uuid.NewV4()

		if err != nil {
			a.log.Println("error")
			return
		}

		// Set the context with the required values to
		// process the request.
		v := Values{
			TraceID: u.String(),
			Now:     time.Now(),
		}
		ctx = context.WithValue(ctx, KeyValues, &v)
		if err := handler(ctx, a.log, w, r, params); err != nil {
			RespondError(ctx, a.log, w, err, http.StatusInternalServerError)
		}
	}

	// Add this handler for the specified verb and route.
	a.TreeMux.Handle(verb, path, h)
}
