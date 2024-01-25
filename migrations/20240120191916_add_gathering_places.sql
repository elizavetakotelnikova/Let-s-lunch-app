-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create type cuisine_type as enum
    (
        'fast_food',
        'russian',
        'italian',
        'eastern',
        'other'
        );

CREATE TABLE  gathering_places (
                           id UUID PRIMARY KEY NOT NULL,
                           country VARCHAR(255) NOT NULL,
                           city VARCHAR(255) NOT NULL,
                           street_name VARCHAR(255) NOT NULL,
                           house_number VARCHAR(255) NOT NULL,
                           building_number int,
                           average_price int,
                           cuisine_type cuisine_type,
                           rating int,
                           phone_number VARCHAR(25)

);
--тут вопрос про opening_hours, про адрес тоже (как хранить)
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
