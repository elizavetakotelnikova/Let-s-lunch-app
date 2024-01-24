package meeting

type MeetingState int

const (
	Active MeetingState = iota
	Cancelled
	Archived
)
