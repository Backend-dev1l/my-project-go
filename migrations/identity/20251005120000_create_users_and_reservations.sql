-- +goose Up
CREATE SCHEMA IF NOT EXISTS identity;

CREATE TABLE IF NOT EXISTS identity.users (
    id UUID PRIMARY KEY,
    balance NUMERIC(20, 2) NOT NULL DEFAULT 0,
    revenue NUMERIC(20, 2) NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS identity.reservations (
    order_id BIGINT PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES identity.users(id),
    amount NUMERIC(20, 2) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_reservations_user_id ON identity.reservations(user_id);

-- +goose Down
DROP INDEX IF EXISTS identity.idx_reservations_user_id;

DROP TABLE IF EXISTS identity.reservations;

DROP TABLE IF EXISTS identity.users;

DROP SCHEMA IF EXISTS identity CASCADE;