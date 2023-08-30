-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN money DECIMAL(10, 2) NOT NULL DEFAULT 5,
    ADD COLUMN total_alc DECIMAL(10, 2) NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
