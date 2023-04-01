-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
	id BIGSERIAL,
	category_id BIGSERIAL,
	name  text NOT NULL,
	description text, 
	price float,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ DEFAULT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY(category_id)
      REFERENCES categories(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products ;
-- +goose StatementEnd
