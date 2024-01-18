package repository

import (
	meeting_domain "cmd/app/entities/meeting"
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type MeetingRepository struct {
	connection *sql.Conn
}

func (m MeetingRepository) FindByID(ctx context.Context, id uuid.UUID) (*meeting_domain.Meeting, error) {
	//TODO implement me
	panic("implement me")
}

func NewMeetingRepository(_ context.Context, connection *sql.Conn) *MeetingRepository {
	return &MeetingRepository{connection: connection}
}
