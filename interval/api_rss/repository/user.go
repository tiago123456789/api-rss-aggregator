package repository

import (
	"database/sql"
	"fmt"

	"github.com/tiago123456789/api-rss-aggregator/interval/api_rss/model"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (userRepository *UserRepository) Create(newUser model.User) (model.User, error) {

	err := userRepository.db.QueryRow(
		`
		INSERT INTO users(name) VALUES ($1) RETURNING id
	`,
		newUser.Name,
	).Scan(&newUser.ID)

	if err != nil {
		return model.User{}, err
	}

	return userRepository.GetById(newUser.ID)
}

func (userRepository *UserRepository) GetById(id int64) (model.User, error) {

	var user model.User
	err := userRepository.db.QueryRow(
		`
		SELECT id, created_at, updated_at, name, api_key FROM users where id = $1 
	`,
		id,
	).Scan(
		&user.ID, &user.CreatedAt,
		&user.UpdatedAt, &user.Name,
		&user.ApiKey,
	)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (userRepository *UserRepository) GetByApiKey(apiKey string) (model.User, error) {
	fmt.Println(apiKey)
	var user model.User
	err := userRepository.db.QueryRow("SELECT id, created_at, updated_at, name, api_key FROM users where api_key = $1",
		&apiKey,
	).Scan(
		&user.ID, &user.CreatedAt,
		&user.UpdatedAt, &user.Name,
		&user.ApiKey,
	)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
