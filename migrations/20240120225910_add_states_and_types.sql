-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO meeting_state(id, state_description)
VALUES (0, 'Active'), (1, 'Cancelled'), (2, 'Archived');

INSERT INTO cuisine_type(id, type_description)
VALUES (0, 'Not Specified'), (1, 'FastFood'), (2, 'Russian'), (3, 'Eastern'), (4, 'Other');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd