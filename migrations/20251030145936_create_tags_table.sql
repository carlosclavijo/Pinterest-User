-- +goose Up
CREATE TABLE tags
(
    id         UUID PRIMARY KEY,
    name       VARCHAR(30) NOT NULL UNIQUE,
    created_at TIMESTAMP   NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE pins_tags
(
    pin_id UUID NOT NULL REFERENCES pins (id) ON DELETE CASCADE,
    tag_id UUID NOT NULL REFERENCES tags (id) ON DELETE CASCADE,
    PRIMARY KEY (pin_id, tag_id)
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE tags;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
