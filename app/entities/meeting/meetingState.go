package meeting

type MeetingState int

const (
	Active = iota + 1
	Cancelled
	Archived
)
