-- +goose Up
-- +goose StatementBegin
create table users
(
    id         uuid               default gen_random_uuid(),
    login      varchar   not null,
    password   varchar   not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),

    primary key (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
