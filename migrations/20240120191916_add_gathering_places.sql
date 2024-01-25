-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE cuisine_type (
    id int PRIMARY KEY NOT NULL,
    type_description VARCHAR(20) NOT NULL
);

CREATE TABLE  gathering_places (
                           id UUID PRIMARY KEY NOT NULL,
                           country VARCHAR(255) NOT NULL,
                           city VARCHAR(255) NOT NULL,
                           street_name VARCHAR(255) NOT NULL,
                           house_number VARCHAR(255) NOT NULL,
                           building_number int,
                           average_price int,
                           cuisine_type int REFERENCES cuisine_type(id),
                           rating int,
                           phone_number VARCHAR(25)
);
--тут вопрос про opening_hours, про адрес тоже (как хранить)
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE gathering_places;
DROP TABLE cuisine_type;