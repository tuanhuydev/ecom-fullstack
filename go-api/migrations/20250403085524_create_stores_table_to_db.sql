-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stores (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,
  CONSTRAINT stores_name_unique UNIQUE (name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stores DROP CONSTRAINT IF EXISTS stores_name_unique;
DROP TABLE IF EXISTS stores;
-- +goose StatementEnd
