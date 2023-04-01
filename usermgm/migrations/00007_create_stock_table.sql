-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stocks (
	id BIGSERIAL,
	product_id BIGSERIAL,
	quantity int,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	

	PRIMARY KEY(id),
	FOREIGN KEY(product_id)
      REFERENCES products(id),
	UNIQUE(product_id)  
  
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stocks ;
-- +goose StatementEnd
