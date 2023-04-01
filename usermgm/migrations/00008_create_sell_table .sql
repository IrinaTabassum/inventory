-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sells (
	id BIGSERIAL,
	customer_id BIGSERIAL,
	total_price float,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	
	PRIMARY KEY(id),
	FOREIGN KEY(customer_id)
      REFERENCES customers(id)
	 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sells ;
-- +goose StatementEnd
