package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/hitesh-goel/loomx/internal/platform/web"
)

// Health provides support for orchestration health checks.
type Health struct {
}

// Check validates the service is healthy and ready to accept requests.
func (h *Health) Check(ctx context.Context, log *log.Logger, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	web.Respond(ctx, log, w, status, http.StatusOK)
	return nil
}
