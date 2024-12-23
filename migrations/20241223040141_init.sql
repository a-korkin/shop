-- +goose Up
-- +goose StatementBegin
create table if not exists public.items (
    id serial primary key,
    title varchar(255),
    price numeric(10, 2)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists public.items;
-- +goose StatementEnd
