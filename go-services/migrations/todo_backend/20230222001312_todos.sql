-- migrate:up

CREATE TABLE todos (
    id bigserial NOT NULL PRIMARY KEY,
    title varchar(255) NOT NULL,
    description text NOT NULL,
    is_completed boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);

-- migrate:down

DROP TABLE todos;
