-- +goose Up
CREATE TABLE boards
(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(500),
    visibility BOOLEAN NOT NULL,
    pin_count INT,
    portrait VARCHAR(200),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE boards;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
