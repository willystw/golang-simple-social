CREATE TABLE IF NOT EXISTS posts (
    id bigserial primary key,
    title text not null,
    user_id bigint not null,
    content text not null,
    created_at TIMESTAMP with time zone NOT NULL DEFAULT NOW()
);

