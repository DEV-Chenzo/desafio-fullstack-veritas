-- +goose Up

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    title VARCHAR(255) NOT NULL,

    description TEXT NOT NULL DEFAULT '',

    status VARCHAR(20) NOT NULL DEFAULT 'TODO',

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    CONSTRAINT tasks_status_check
        CHECK (status IN ('TODO', 'IN_PROGRESS', 'DONE'))
);

-- +goose Down

DROP TABLE IF EXISTS tasks;

DROP EXTENSION IF EXISTS "pgcrypto";