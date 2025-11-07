-- +goose Up
create table users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login     text not null,
    password  text not null
);

-- +goose Down
drop table users;