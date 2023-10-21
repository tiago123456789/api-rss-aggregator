package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/repository"
	"github.com/tiago123456789/api-rss-aggregator/pkg/response_http"
)

type PostController struct {
	repository repository.PostRepository
}

func NewPostRepository(db *sql.DB) *PostController {
	return &PostController{
		repository: *repository.NewPostepository(db),
	}
}

func (controller *PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id")
	userIdString := fmt.Sprintf("%d", userId)
	userIntParsed, _ := strconv.ParseInt(userIdString, 10, 64)
	posts, err := controller.repository.GetPostsByUserId(userIntParsed)
	if err != nil {
		log.Printf("Error => %v", err)
	}

	response_http.ReturnJson(w, 200, posts)
}
