-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id                 BIGSERIAL    NOT NULL PRIMARY KEY,
    email              VARCHAR(255) NOT NULL UNIQUE,
    encrypted_password VARCHAR(255) NOT NULL,
    created_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
