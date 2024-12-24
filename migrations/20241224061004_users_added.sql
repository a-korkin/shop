-- +goose Up
-- +goose StatementBegin
create table if not exists public.users (
    id serial primary key,
    last_name varchar(255),
    first_name varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists public.users;
-- +goose StatementEnd
