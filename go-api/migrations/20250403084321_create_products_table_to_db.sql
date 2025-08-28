-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    price DECIMAL NOT NULL,
    description VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    thumbnail VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
  )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
  DROP TABLE IF EXISTS products;
-- +goose StatementEnd
