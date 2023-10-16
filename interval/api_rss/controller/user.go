package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tiago123456789/api-rss-aggregator/interval/api_rss/model"
	"github.com/tiago123456789/api-rss-aggregator/interval/api_rss/repository"
	"github.com/tiago123456789/api-rss-aggregator/pkg/response_http"
)

type UserController struct {
	repository repository.UserRepository
}

func New(db *sql.DB) *UserController {
	return &UserController{
		repository: *repository.New(db),
	}
}

func (controller *UserController) GetRepository() *repository.UserRepository {
	return &controller.repository
}

func (controller *UserController) SayHi(w http.ResponseWriter, r *http.Request) {
	type hiMessage struct {
		Message string `json:"message"`
	}

	fmt.Printf("Context value here => %s", r.Context().Value("user_id"))

	response_http.ReturnJson(w, 200, hiMessage{
		Message: "Hi my friend",
	})
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	type hiMessage struct {
		Message string `json:"message"`
	}

	var newUser model.User
	json.NewDecoder(r.Body).Decode(&newUser)

	userCreated, err := controller.repository.Create(newUser)
	if err != nil {
		log.Printf("Error => %v", err)
	}

	response_http.ReturnJson(w, 201, userCreated)
}

func (controller *UserController) Healthz(w http.ResponseWriter, r *http.Request) {
	response_http.Return200AndEmpty(w)
}

func (controller *UserController) Return400(w http.ResponseWriter, r *http.Request) {
	erros := []string{"The name is required", "The email is required"}
	response_http.Return400(w, response_http.ErrorMessage{
		StatusCode: 400,
		Error:      erros,
	})
}
