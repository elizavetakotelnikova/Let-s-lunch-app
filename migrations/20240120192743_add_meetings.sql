-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create type meeting_state as enum
    (
        'active',
        'cancelled',
        'archived'
        );

CREATE TABLE  meetings (
                  id UUID PRIMARY KEY NOT NULL,
                  gathering_place_id UUID REFERENCES gathering_places(id),
                  initiators_id  UUID REFERENCES users(id),
                  time_start     timestamp,
                  time_end       timestamp,
                  max_participants  int,
                  state        meeting_state
);
CREATE TABLE meetings_history (
                                  user_id UUID REFERENCES users(id),
                                  meeting_id UUID REFERENCES meetings(id)
);
ALTER TABLE users
ADD current_meeting_id UUID REFERENCES meetings(id);
--тут вопрос про registration_end, в доке он стоит со знаком "?"
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE meetings;
DROP TYPE meeting_state;
DROP TABLE meetings_history;
ALTER TABLE users
DROP COLUMN current_meeting_id;
-- +goose StatementEnd
