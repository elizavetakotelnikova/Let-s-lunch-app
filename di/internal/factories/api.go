package factories

import (
	gathering_place_api "cmd/app/entities/gatheringPlace/api"
	meeting_api "cmd/app/entities/meeting/api"
	user_api "cmd/app/entities/user/api"
	"cmd/di/internal/lookup"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateRouter(ctx context.Context, c lookup.Container) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/meetings", func(r chi.Router) {
			r.Route("/{meetingID}", func(r chi.Router) {
				r.Get("/", c.API().FindMeetingHandler(ctx).ServeHTTP)
			})
		})

		r.Route("/user", func(r chi.Router) {
			r.Route("/create", func(r chi.Router) {
				r.Post("/", c.API().CreateUser(ctx).ServeHTTP)
			})

			r.Route("/find_by_id", func(r chi.Router) {
				r.Route("/{userID}", func(r chi.Router) {
					r.Get("/", c.API().FindUserHandler(ctx).ServeHTTP)
				})
			})
		})

		r.Route("/places", func(r chi.Router) {
			r.Route("/{placeID}", func(r chi.Router) {
				r.Get("/", c.API().FindGatheringPlaceHandler(ctx).ServeHTTP)
			})
		})
	})

	return r
}

func CreateAPIFindMeetingHandler(ctx context.Context, c lookup.Container) *meeting_api.FindMeetingByIdHandler {
	return meeting_api.NewFindMeetingByIdHandler(
		c.UseCases().FindMeeting(ctx),
	)
}

func CreateAPIFindUserHandler(ctx context.Context, c lookup.Container) *user_api.FindUserByIdHandler {
	return user_api.NewFindUserByIdHandler(
		c.UseCases().FindUser(ctx),
	)
}

func CreateAPIFindGatheringPlaceHandler(ctx context.Context, c lookup.Container) *gathering_place_api.FindGatheringPlaceByIdHandler {
	return gathering_place_api.NewFindGatheringCaseByIdHandler(
		c.UseCases().FindGatheringPlace(ctx),
	)
}

func CreateAPICreateUser(ctx context.Context, c lookup.Container) *user_api.CreateUserHandler {
	return user_api.NewCreateUserHandler(
		c.UseCases().CreateUser(ctx),
	)
}
