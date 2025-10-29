-- +goose Up
CREATE TABLE users
(
    id            UUID PRIMARY KEY,
    first_name    VARCHAR(100) NOT NULL,
    last_name     VARCHAR(100) NOT NULL,
    user_name     VARCHAR(30)  NOT NULL UNIQUE,
    email         VARCHAR(320) NOT NULL UNIQUE,
    password      VARCHAR(255) NOT NULL,
    gender        CHAR(1)      NOT NULL DEFAULT 'O' CHECK (gender IN ('O', 'M', 'F')),
    birth_date    DATE         NOT NULL,
    country       CHAR(2)      NOT NULL,
    language      CHAR(2)      NOT NULL,
    phone         VARCHAR(16) UNIQUE,
    information   VARCHAR(500),
    profile_pic   VARCHAR(255),
    web_site      VARCHAR(255),
    visibility    BOOL                  DEFAULT TRUE,
    last_login_at TIMESTAMP         NOT NULL,
    created_at    TIMESTAMP         NOT NULL,
    updated_at    TIMESTAMP         NOT NULL,
    deleted_at    TIMESTAMP
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE users;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

-- goose postgres "user=postgres password=abc12345 dbname=pinterest-user sslmode=disable" down
-- goose postgres "user=postgres password=abc12345 dbname=pinterest-user sslmode=disable" up