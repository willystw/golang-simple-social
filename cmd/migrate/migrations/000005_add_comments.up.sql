CREATE TABLE IF NOT EXISTS comments (
    id bigserial primary key,
    post_id bigint not null references posts(id),
    user_id bigint not null references users(id),
    content text,
    created_at TIMESTAMP with time zone NOT NULL DEFAULT NOW()
);