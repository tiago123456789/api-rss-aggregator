-- +goose Up
-- +goose StatementBegin
CREATE TABLE feeds_follows(
    user_id INT,
    feed_id INT,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feeds_follows
-- +goose StatementEnd
