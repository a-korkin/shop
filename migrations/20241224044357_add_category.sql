-- +goose Up
-- +goose StatementBegin
alter table if exists public.items 
    add column if not exists category varchar(255) default '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table if exists public.items drop column if exists category; 
-- +goose StatementEnd
