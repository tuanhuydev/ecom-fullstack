-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS collections (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT collections_name_unique UNIQUE (name)
);
INSERT INTO collections (name, description) VALUES ('Default Collection', 'This is the default collection.');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE collections DROP CONSTRAINT IF EXISTS stores_name_unique;
DROP TABLE IF EXISTS collections;
-- +goose StatementEnd
