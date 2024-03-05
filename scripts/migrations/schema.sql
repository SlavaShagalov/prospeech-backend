CREATE TABLE IF NOT EXISTS users
(
    id              serial    NOT NULL PRIMARY KEY,
    username        text      NOT NULL UNIQUE,
    hashed_password varchar   NOT NULL,
    email           varchar   NOT NULL,
    name            varchar   NOT NULL,
    avatar          varchar   NULL,
    created_at      timestamp NOT NULL DEFAULT now(),
    updated_at      timestamp NOT NULL DEFAULT now()
);
