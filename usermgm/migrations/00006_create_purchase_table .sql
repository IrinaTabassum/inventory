-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS purchases (
	id BIGSERIAL,
	supplier_id BIGSERIAL,
	product_id BIGSERIAL,
	quantity int,
	unit_price float,
	total_price float,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	
	PRIMARY KEY(id),
	FOREIGN KEY(supplier_id )
      REFERENCES suppliers(id),
	FOREIGN KEY(product_id)
      REFERENCES products(id)  
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS purchases;
-- +goose StatementEnd
