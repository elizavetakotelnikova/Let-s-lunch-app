package definitions

import (
	meeting_domain "cmd/app/entities/meeting"
	meeting_api "cmd/app/entities/meeting/api"
	meeting_usecase "cmd/app/entities/meeting/usecases"
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
	FindMeetingHandler *meeting_api.FindMeeting
}

type UseCaseContainer struct {
	FindMeeting *meeting_usecase.FindMeeting
}

type RepositoryContainer struct {
	meetingRepository meeting_domain.MeetingRepository `di:"set"`
}
