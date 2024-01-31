package factories

import (
	meeting_api "cmd/app/entities/meeting/api"
	"cmd/di/internal/lookup"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateRouter(ctx context.Context, c lookup.Container) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/meetings", func(r chi.Router) {
		r.Route("/{meetingID}", func(r chi.Router) {
			r.Get("/", c.API().FindMeetingHandler(ctx).ServeHTTP)
		})
	})

	return r
}

func CreateAPIFindMeetingHandler(ctx context.Context, c lookup.Container) *meeting_api.FindMeeting {
	return meeting_api.NewFindMeeting(
		c.UseCases().FindMeeting(ctx),
	)
}
