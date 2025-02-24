create EXTENSION if not exists citext;

CREATE TABLE IF NOT EXISTS users (
    id bigserial primary key,
    username varchar(255) unique not null,
    email citext unique not null,
    password bytea not null,
    created_at TIMESTAMP with time zone NOT NULL DEFAULT NOW()
);

