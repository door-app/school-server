
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users
  (id BIGINT, name VARCHAR(255));

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
