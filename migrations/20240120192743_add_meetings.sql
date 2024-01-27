-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE meeting_state (
    id int PRIMARY KEY NOT NULL,
    state_description VARCHAR(20) NOT NULL
);

CREATE TABLE  meetings (
                  id UUID PRIMARY KEY NOT NULL,
                  gathering_place_id UUID REFERENCES gathering_places(id),
                  initiators_id  UUID REFERENCES users(id),
                  time_start     timestamp,
                  time_end       timestamp,
                  max_participants  int,
                  state        int REFERENCES meeting_state(id)
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
ALTER TABLE users
DROP COLUMN current_meeting_id;
DROP TABLE meetings_history;
DROP TABLE meetings;
DROP TABLE meeting_state;
-- +goose StatementEnd
