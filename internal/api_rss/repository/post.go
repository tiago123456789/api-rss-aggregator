package repository

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/model"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (repository *PostRepository) GetPostsByUserId(userId int64) ([]model.Post, error) {
	rows, err := repository.db.Query(
		`
			select 
				p.id, p.title, p.link, p.published_at,
				p.created_at, p.updated_at, p.feed_id
			from posts as p
			inner join feeds as f on (f.id = p.feed_id)
			inner join feeds_follows as ff on (ff.feed_id = f.id)
			where ff.user_id = $1
		`,
		userId,
	)

	if err != nil {
		return []model.Post{}, err
	}

	posts := []model.Post{}
	for rows.Next() {
		post := model.Post{}
		err = rows.Scan(
			&post.ID, &post.Title, &post.Link, &post.PublishedAt,
			&post.CreatedAt, &post.UpdatedAt, &post.FeedId,
		)
		if err != nil {
			return []model.Post{}, err
		}
		posts = append(posts, post)
	}
	rows.Close()

	if err != nil {
		return []model.Post{}, err
	}

	return posts, nil
}

func (repository *PostRepository) GetPostByFeedIds(feedIs []int64) ([]model.Post, error) {
	var feedIsParam []string

	for _, value := range feedIs {
		valueString := strconv.FormatInt(int64(value), 10)
		feedIsParam = append(feedIsParam, valueString)
	}

	rows, err := repository.db.Query(
		`
		SELECT Title, feed_id FROM posts where feed_id = ANY($1::int[])
	`,
		"{"+strings.Join(feedIsParam, ",")+"}",
	)

	if err != nil {
		return []model.Post{}, err
	}

	posts := []model.Post{}
	for rows.Next() {
		post := model.Post{}
		err = rows.Scan(
			&post.Title, &post.FeedId,
		)
		if err != nil {
			log.Println(err)
		}
		posts = append(posts, post)
	}
	rows.Close()

	if err != nil {
		return []model.Post{}, err
	}

	return posts, nil
}

func (repository *PostRepository) InsertMany(newPosts []model.Post) error {
	txn, err := repository.db.Begin()
	if err != nil {
		return err
	}

	stmt, _ := txn.Prepare(pq.CopyIn(
		"posts",
		"id", "created_at", "updated_at",
		"title", "link", "published_at",
		"feed_id",
	))

	for _, post := range newPosts {
		_, err := stmt.Exec(
			uuid.NewString(), time.Now(), time.Now(),
			post.Title, post.Link, post.PublishedAt,
			post.FeedId,
		)
		if err != nil {
			return err
		}
	}

	err = stmt.Close()
	if err != nil {
		return err
	}
	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}
