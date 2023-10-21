package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/model"
)

type FeedRepository struct {
	db *sql.DB
}

func NewFeedRepository(db *sql.DB) *FeedRepository {
	return &FeedRepository{
		db: db,
	}
}

func (repository *FeedRepository) SetLastFetchedAt(feedId int64) error {
	_, err := repository.db.Exec(
		`
		UPDATE feeds set last_fetched_at = CURRENT_TIMESTAMP where 
		id = $1
	`,
		feedId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repository *FeedRepository) Create(newFeed model.Feed) (model.Feed, error) {
	err := repository.db.QueryRow(
		`
		INSERT INTO feeds(name, created_at, updated_at, url, user_id) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id
	`,
		newFeed.Name, time.Now(), time.Now(), newFeed.Url, newFeed.UserID,
	).Scan(&newFeed.ID)

	if err != nil {
		return model.Feed{}, err
	}

	return newFeed, nil
}

func (repository *FeedRepository) GetFeedsToProcessRSS() ([]model.Feed, error) {
	rows, err := repository.db.Query(
		`
		SELECT id, name, user_id, url, created_at, updated_at from feeds
		WHERE last_fetched_at is null OR CURRENT_TIMESTAMP > last_fetched_at
	`,
	)

	if err != nil {
		return []model.Feed{}, err
	}

	feeds := []model.Feed{}
	for rows.Next() {
		feed := model.Feed{}
		err = rows.Scan(
			&feed.ID, &feed.Name, &feed.UserID,
			&feed.Url, &feed.CreatedAt, &feed.UpdatedAt,
		)
		if err != nil {
			return []model.Feed{}, err
		}
		feeds = append(feeds, feed)
	}
	rows.Close()

	if err != nil {
		return []model.Feed{}, err
	}

	return feeds, nil
}

func (repository *FeedRepository) GetFeeds() ([]model.Feed, error) {
	rows, err := repository.db.Query(
		`
		SELECT id, name, user_id, url, created_at, updated_at from feeds
	`,
	)
	feeds := []model.Feed{}
	for rows.Next() {
		feed := model.Feed{}
		err = rows.Scan(
			&feed.ID, &feed.Name, &feed.UserID,
			&feed.Url, &feed.CreatedAt, &feed.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
		}
		feeds = append(feeds, feed)
	}
	rows.Close()

	if err != nil {
		return []model.Feed{}, err
	}

	return feeds, nil
}

func (repository *FeedRepository) Follow(newFeed model.Feed) (model.Feed, error) {
	_, err := repository.db.Query(
		`
		INSERT INTO feeds_follows(user_id, feed_id) 
		VALUES ($1, $2)
	`,
		newFeed.UserID, newFeed.ID,
	)

	if err != nil {
		return model.Feed{}, err
	}

	return newFeed, nil
}

func (repository *FeedRepository) Unfollow(newFeed model.Feed) (model.Feed, error) {
	_, err := repository.db.Query(
		`
		DELETE FROM feeds_follows where user_id = $1 and feed_id = $2
	`,
		newFeed.UserID, newFeed.ID,
	)

	if err != nil {
		return model.Feed{}, err
	}

	return newFeed, nil
}
