-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO meeting_state(id, state_description)
VALUES (1, 'Active'), (2, 'Cancelled'), (3, 'Archived');

INSERT INTO cuisine_type(id, type_description)
VALUES (1, 'FastFood'), (2, 'Russian'), (3, 'Eastern'), (4, 'Other');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DELETE FROM cuisine_type WHERE id >= 0 AND id <= 3;
DELETE FROM meeting_state WHERE id >= 0 AND id <= 2;