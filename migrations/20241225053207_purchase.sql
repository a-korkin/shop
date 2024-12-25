-- +goose Up
-- +goose StatementBegin
create table if not exists public.purchases 
(
    id serial primary key,
    user_id integer not null references public.users(id),
    item_id integer not null references public.items(id),
    time_of_purchase timestamp not null default now(),
    count_items integer not null default 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists public.purchases;
-- +goose StatementEnd
