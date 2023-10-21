package model

import "time"

type User struct {
	ID        int64     `json:"Id"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}
