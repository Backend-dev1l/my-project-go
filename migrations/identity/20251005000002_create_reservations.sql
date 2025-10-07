-- +goose Up
CREATE TABLE IF NOT EXISTS identity.reservations (
    order_id BIGINT PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES identity.users(id),
    amount NUMERIC(20, 2) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS identity.reservations;