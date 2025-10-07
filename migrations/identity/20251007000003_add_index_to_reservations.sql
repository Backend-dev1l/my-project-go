-- +goose Up
CREATE INDEX IF NOT EXISTS idx_reservations_user_id ON identity.reservations(user_id);

-- +goose Down
DROP INDEX IF EXISTS identity.idx_reservations_user_id;