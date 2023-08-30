-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS battle_records
(
    id            BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    attackerId    BIGINT  NOT NULL,
    attackerPower DECIMAL(10, 2)   NOT NULL,
    defenderId    BIGINT  NOT NULL,
    defenderPower DECIMAL(10, 2)   NOT NULL,
    winner        TINYINT NOT NULL,
    FOREIGN KEY (attackerId) REFERENCES users (id),
    FOREIGN KEY (defenderId) REFERENCES users (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE battle_records
-- +goose StatementEnd
