-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS email_schedule
(
    id                BIGSERIAL    NOT NULL PRIMARY KEY,
    email_template_id INT REFERENCES email_templates (id),
    recipients        TEXT,
    execute_after     TIMESTAMP,
    execution_period  INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS email_schedule;
-- +goose StatementEnd
