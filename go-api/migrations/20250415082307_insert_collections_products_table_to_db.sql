-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS collections_products (
  collection_id UUID NOT NULL,
  product_id UUID NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,
  PRIMARY KEY (collection_id, product_id), -- Composite primary key
  CONSTRAINT collections_products_collection_id_fkey FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE,
  CONSTRAINT collections_products_product_id_fkey FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE collections_products DROP CONSTRAINT IF EXISTS collections_products_collection_id_fkey;
ALTER TABLE collections_products DROP CONSTRAINT IF EXISTS collections_products_product_id_fkey;
DROP TABLE IF EXISTS collections_products;
-- +goose StatementEnd
