package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/model"
	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/repository"
	"github.com/tiago123456789/api-rss-aggregator/pkg/response_http"
)

type UserController struct {
	repository repository.UserRepository
}

func NewUsercontroller(db *sql.DB) *UserController {
	return &UserController{
		repository: *repository.NewUserRepository(db),
	}
}

func (controller *UserController) GetRepository() *repository.UserRepository {
	return &controller.repository
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	json.NewDecoder(r.Body).Decode(&newUser)

	if len(newUser.Name) == 0 {
		response_http.Return400(w, response_http.ErrorMessage{
			StatusCode: 400,
			Error:      []string{"The field name is required"},
		})
		return
	}

	userCreated, err := controller.repository.Create(newUser)
	if err != nil {
		log.Printf("Error => %v", err)
	}

	response_http.ReturnJson(w, 201, userCreated)
}

func (controller *UserController) Healthz(w http.ResponseWriter, r *http.Request) {
	response_http.Return200AndEmpty(w)
}
