package definitions

import (
	gathering_place_api "cmd/app/entities/gatheringPlace/api"
	gathering_place_repository "cmd/app/entities/gatheringPlace/repository"
	gathering_place_usecase "cmd/app/entities/gatheringPlace/usecases"
	meeting_api "cmd/app/entities/meeting/api"
	meeting_repository "cmd/app/entities/meeting/repository"
	meeting_usecase "cmd/app/entities/meeting/usecases"
	user_api "cmd/app/entities/user/api"
	user_repository "cmd/app/entities/user/repository"
	user_usecase "cmd/app/entities/user/usecases"
	chi "github.com/go-chi/chi/v5"

	"cmd/app/config"
	"database/sql"
	"log"
	"net/http"
)

// Container is a root dependency injection container. It is required to describe
// your services.
type Container struct {
	Config config.Params `di:"required"`
	Logger *log.Logger
	DB     *sql.DB `di:"close"`

	Server *http.Server `di:"public,close" factory-file:"server"`
	Router *chi.Mux     `factory-file:"api"`

	API          APIContainer
	UseCases     UseCaseContainer
	Repositories RepositoryContainer
}

type APIContainer struct {
	CreateUserHandler           *user_api.CreateUserHandler
	UpdateUserHandler           *user_api.UpdateUserHandler
	DeleteUserHandler           *user_api.DeleteUserHandler
	FindUserHandler             *user_api.FindUserByIdHandler
	FindMeetingHandler          *meeting_api.FindMeetingByIdHandler
	CreateMeetingHandler        *meeting_api.CreateMeetingHandler
	FindGatheringPlaceHandler   *gathering_place_api.FindGatheringPlaceByIdHandler
	CreateGatheringPlaceHandler *gathering_place_api.CreateGatheringPlaceHandler
	UpdateGatheringPlaceHandler *gathering_place_api.UpdateGatheringPlaceHandler
	DeleteGatheringPlaceHandler *gathering_place_api.DeleteGatheringPlaceHandler
}

type UseCaseContainer struct {
	FindUser             *user_usecase.FindUserByIdUseCase
	CreateUser           *user_usecase.CreateUserUseCase
	UpdateUser           *user_usecase.UpdateUserUseCase
	DeleteUser           *user_usecase.DeleteUserUseCase
	FindMeeting          *meeting_usecase.FindMeetingByIdUseCase
	CreateMeeting        *meeting_usecase.CreateMeetingUseCase
	FindGatheringPlace   *gathering_place_usecase.FindGatheringPlaceByIdUseCase
	CreateGatheringPlace *gathering_place_usecase.CreateGatheringPlaceUseCase
	UpdateGatheringPlace *gathering_place_usecase.UpdateGatheringPlaceUseCase
	DeleteGatheringPlace *gathering_place_usecase.DeleteGatheringPlaceUseCase
}

type RepositoryContainer struct {
	meetingRepository        meeting_repository.MeetingsRepository       `di:"set"`
	userRepository           user_repository.UsersRepository             `di:"set"`
	gatheringPlaceRepository gathering_place_repository.PlacesRepository `di:"set"`
}
