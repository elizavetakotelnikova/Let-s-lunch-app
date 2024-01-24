package meeting

import (
	"context"
	"github.com/google/uuid"
)

type MeetingRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*Meeting, error)
}
