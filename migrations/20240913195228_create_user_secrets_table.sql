-- +goose Up
-- +goose StatementBegin
create table user_secrets
(
    id         uuid               default gen_random_uuid(),
    user_id    uuid      not null references users (id),
    type       varchar   not null,
    name       varchar   not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),

    primary key (id)
);

create index user_secrets_user_created_at_inx on user_secrets (user_id, created_at)


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_secrets;
-- +goose StatementEnd
