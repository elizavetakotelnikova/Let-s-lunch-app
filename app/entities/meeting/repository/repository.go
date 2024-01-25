package repository

import (
	meeting_domain "cmd/app/entities/meeting"
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type MeetingRepository struct {
	connection *sql.DB
}

func (m MeetingRepository) FindByID(ctx context.Context, id uuid.UUID) (*meeting_domain.Meeting, error) {
	panic("implement me")
	m.connection.ExecContext()
}

func NewMeetingRepository(_ context.Context, connection *sql.DB) *MeetingRepository {
	return &MeetingRepository{connection: connection}
}
