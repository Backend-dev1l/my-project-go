
-- +goose Up
CREATE SCHEMA IF NOT EXISTS identity;

CREATE TABLE IF NOT EXISTS identity.users (
    id UUID PRIMARY KEY,
    balance NUMERIC(20, 2) NOT NULL DEFAULT 0,
    revenue NUMERIC(20, 2) NOT NULL DEFAULT 0
);

-- +goose Down
DROP TABLE IF EXISTS identity.users;