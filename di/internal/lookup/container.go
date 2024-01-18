// Code generated by DIGEN; DO NOT EDIT.
// This file was generated by Dependency Injection Container Generator 0.1.0 (built at 2023-10-21T19:51:59Z).
// See docs at https://github.com/strider2038/digen

package lookup

import (
	"cmd/app/config"
	meeting_domain "cmd/app/entities/meeting"
	meeting_api "cmd/app/entities/meeting/api"
	meeting_usecase "cmd/app/entities/meeting/usecases"
	"context"
	"database/sql"
	"log"
	"net/http"
)

type Container interface {
	// SetError sets the first error into container. The error is used in the public container to return an initialization error.
	SetError(err error)

	Config(ctx context.Context) config.Params
	Logger(ctx context.Context) *log.Logger
	DB(ctx context.Context) *sql.DB
	Server(ctx context.Context) *http.Server

	API() APIContainer
	UseCases() UseCaseContainer
	Repositories() RepositoryContainer
}

type APIContainer interface {
	FindMeetingHandler(ctx context.Context) *meeting_api.FindMeeting
}

type UseCaseContainer interface {
	FindMeeting(ctx context.Context) *meeting_usecase.FindMeeting
}

type RepositoryContainer interface {
	MeetingRepository(ctx context.Context) meeting_domain.MeetingRepository
}
