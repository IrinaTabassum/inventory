-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS suppliers (
	id BIGSERIAL,
	name  VARCHAR(20) NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ DEFAULT NULL,

	PRIMARY KEY(id),
	UNIQUE(name)

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS suppliers ;
-- +goose StatementEnd
