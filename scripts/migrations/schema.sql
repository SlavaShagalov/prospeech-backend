CREATE TABLE IF NOT EXISTS users
(
    id                      bigserial NOT NULL PRIMARY KEY,
    username                text      NOT NULL UNIQUE,
    hashed_password         varchar   NOT NULL,
    email                   varchar   NOT NULL,
    name                    varchar   NOT NULL,
    avatar                  varchar   NULL,
    created_at              timestamp NOT NULL DEFAULT now(),
    updated_at              timestamp NOT NULL DEFAULT now(),
    untitled_speeches_count int       NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS audios
(
    id            bigserial NOT NULL PRIMARY KEY,
    user_id       bigint    NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    title         varchar   NOT NULL DEFAULT '',
    url           varchar   NOT NULL,
    created_at    timestamp NOT NULL DEFAULT now(),
    updated_at    timestamp NOT NULL DEFAULT now(),
    -- ML data
    text          text      NULL,
    words         text[]    NULL,
    start_times   float[]   NULL,
    end_times     float[]   NULL,
    words_per_min int       NULL,
    duration      interval  NULL
);

CREATE TABLE IF NOT EXISTS videos
(
    id         bigserial NOT NULL PRIMARY KEY,
    user_id    bigint    NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    title      varchar   NOT NULL DEFAULT '',
    url        varchar   NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
);
