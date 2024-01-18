package factories

import (
	meeting_api "cmd/app/entities/meeting/api"
	"cmd/di/internal/lookup"
	"context"
	"net/http"
)

func CreateRouter(ctx context.Context, c lookup.Container) *http.ServeMux {

}

func CreateAPIFindMeetingHandler(ctx context.Context, c lookup.Container) *meeting_api.FindMeeting {
	panic("not implemented")
}
