package model

import "time"

type Feed struct {
	ID            int64     `json:"Id"`
	CreatedAt     time.Time `json:"createAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Name          string    `json:"name"`
	Url           string    `json:"url"`
	UserID        any       `json:"user_id"`
	LastFetchedAt time.Time `json:"last_fetched_at"`
}
