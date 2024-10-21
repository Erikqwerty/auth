-- +goose Up
CREATE TABLE user_log (
    id SERIAL PRIMARY KEY,
    action_type VARCHAR(50) NOT NULL,
    action_details TEXT,
    action_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Время выполнения
);
-- +goose Down
DROP TABLE IF EXISTS user_log;