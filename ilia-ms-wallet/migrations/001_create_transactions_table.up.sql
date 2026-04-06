-- +migrate Up
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    amount NUMERIC(12,2) NOT NULL,
    type VARCHAR(10) NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS transactions;
