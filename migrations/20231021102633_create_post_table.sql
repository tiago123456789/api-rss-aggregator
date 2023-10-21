-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts(
    id uuid PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    title TEXT NOT NULL,
    link TEXT NOT NULL,
    published_at TIMESTAMP NOT NULL,
    feed_id INT,
    FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
