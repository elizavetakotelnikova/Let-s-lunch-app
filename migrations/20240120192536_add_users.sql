-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE  users (
                           id UUID PRIMARY KEY NOT NULL,
                           username VARCHAR(255) NOT NULL,
                           display_name VARCHAR(255) NOT NULL,
                           rating int
                    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE users;
-- +goose StatementEnd
