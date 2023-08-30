-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id     BIGINT AUTO_INCREMENT         NOT NULL PRIMARY KEY,
    tgId   BIGINT         NOT NULL,

    name    VARCHAR(512) NOT NULL,
    power  DECIMAL(10, 2) NOT NULL,
    title  BIGINT         NOT NULL,
    energy DECIMAL(10, 2) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
