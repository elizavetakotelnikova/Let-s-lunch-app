package factories

import (
	"cmd/di/internal/lookup"
	"context"
	"net/http"
)

func CreateServer(ctx context.Context, c lookup.Container) *http.Server {
	return &http.Server{
		Handler: c.API().FindMeetingHandler(ctx),
	}
}
