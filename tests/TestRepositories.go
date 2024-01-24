package tests

import (
	"cmd/app/entities/meeting"
	"cmd/app/entities/meeting/repository"
	"context"
	"testing"
)

// just something (I was trying to check enum)
func TestMeetingsRepositoryPassedCorrectShouldReturnSaved(t *testing.T) {
	var databaseRepository = repository.MeetingsDatabaseRepository{}
	var currentMeeting = meeting.Meeting{}
	currentMeeting.State = meeting.Cancelled
	var ctx = context.Background()
	_, err := databaseRepository.Create(ctx, &currentMeeting)
	if err != nil {
		t.Fatalf("Cannot save user %w", err)
	}
}
