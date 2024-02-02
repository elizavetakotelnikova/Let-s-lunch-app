-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE users
ADD COLUMN phone_number VARCHAR(255);
ALTER TABLE users
ADD COLUMN birthday date;
ALTER TABLE users
DROP column age;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE users
ADD COLUMN age int;

ALTER TABLE users
DROP COLUMN phone_number,
DROP column birthday;
-- +goose StatementEnd
