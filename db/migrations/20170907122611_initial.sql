
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users (
  id            BIGSERIAL PRIMARY KEY NOT NULL,
  created_at    TIMESTAMP NOT NULL,
  updated_at    TIMESTAMP,
  deleted_at    TIMESTAMP,
  email_address VARCHAR
);

CREATE TABLE sessions (
  id            BIGSERIAL PRIMARY KEY NOT NULL,
  created_at    TIMESTAMP NOT NULL,
  deleted_at    TIMESTAMP,
  user_id       BIGSERIAL REFERENCES users (id) NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE sessions;
DROP TABLE users;