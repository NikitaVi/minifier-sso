-- +goose Up
create table user_premium
(
    user_id    UUID PRIMARY KEY REFERENCES users (id) ON DELETE CASCADE,
    is_active  BOOLEAN NOT NULL DEFAULT TRUE,
    start_date TIMESTAMP        DEFAULT now(),
    end_date   TIMESTAMP
);

-- +goose Down
drop table user_premium;
