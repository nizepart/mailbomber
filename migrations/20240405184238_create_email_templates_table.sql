-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS email_templates
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    subject VARCHAR(255),
    body TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS email_templates;
-- +goose StatementEnd
