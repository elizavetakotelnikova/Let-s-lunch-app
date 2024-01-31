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
		r.Route("/meeting", func(r chi.Router) {
			r.Route("/find", func(r chi.Router) {
				r.Get("/{meetingID}", c.API().FindMeetingHandler(ctx).ServeHTTP)

			})
			r.Post("/create", c.API().CreateMeetingHandler(ctx).ServeHTTP)
		})

		r.Route("/user", func(r chi.Router) {
			r.Post("/create", c.API().CreateUserHandler(ctx).ServeHTTP)

			r.Route("/find_by_id", func(r chi.Router) {
				r.Get("/{userID}", c.API().FindUserHandler(ctx).ServeHTTP)
			})

			r.Route("/update", func(r chi.Router) {
				r.Put("/{userID}", c.API().UpdateUserHandler(ctx).ServeHTTP)
				r.Delete("/{userID}", c.API().DeleteUserHandler(ctx).ServeHTTP)
			})
		})

		r.Route("/gatheringPlace", func(r chi.Router) {
			r.Route("/find_by_id", func(r chi.Router) {
				r.Get("/{placeID}", c.API().FindGatheringPlaceHandler(ctx).ServeHTTP)
			})
			r.Post("/create", c.API().CreateGatheringPlaceHandler(ctx).ServeHTTP)
			r.Route("/update", func(r chi.Router) {
				r.Put("/{placeID}", c.API().UpdateGatheringPlaceHandler(ctx).ServeHTTP)
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

func CreateAPICreateUserHandler(ctx context.Context, c lookup.Container) *user_api.CreateUserHandler {
	return user_api.NewCreateUserHandler(
		c.UseCases().CreateUser(ctx),
	)
}

func CreateAPIUpdateUserHandler(ctx context.Context, c lookup.Container) *user_api.UpdateUserHandler {
	return user_api.NewUpdateUserHandler(
		c.UseCases().UpdateUser(ctx),
	)
}

func CreateAPIDeleteUserHandler(ctx context.Context, c lookup.Container) *user_api.DeleteUserHandler {
	return user_api.NewDeleteUserHandler(
		c.UseCases().DeleteUser(ctx),
	)
}

func CreateAPICreateMeetingHandler(ctx context.Context, c lookup.Container) *meeting_api.CreateMeetingHandler {
	return meeting_api.NewCreateMeetingHandler(
		c.UseCases().CreateMeeting(ctx),
	)
}

func CreateAPICreateGatheringPlaceHandler(ctx context.Context, c lookup.Container) *gathering_place_api.CreateGatheringPlaceHandler {
	return gathering_place_api.NewCreateGatheringPlaceHandler(
		c.UseCases().CreateGatheringPlace(ctx),
	)
}

func CreateAPIUpdateGatheringPlaceHandler(ctx context.Context, c lookup.Container) *gathering_place_api.UpdateGatheringPlaceHandler {
	return gathering_place_api.NewUpdateGatheringPlaceHandler(
		c.UseCases().UpdateGatheringPlace(ctx))
}
