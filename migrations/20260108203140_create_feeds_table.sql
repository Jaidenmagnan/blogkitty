-- +goose Up
CREATE TABLE IF NOT EXISTS feeds(
  id INTEGER PRIMARY KEY UNIQUE,
  feed_url VARCHAR(100),
  discord_channel_id VARCHAR(100),
  latest_post_guid VARCHAR(128)
);
-- +goose StatementBegin
SELECT
  'up SQL query';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';
-- +goose StatementEnd
