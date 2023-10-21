package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/model"
	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/repository"
	"github.com/tiago123456789/api-rss-aggregator/pkg/response_http"
)

type FeedController struct {
	repository repository.FeedRepository
}

func NewFeedController(db *sql.DB) *FeedController {
	return &FeedController{
		repository: *repository.NewFeedRepository(db),
	}
}

func (controller *FeedController) Create(w http.ResponseWriter, r *http.Request) {
	var newFeed model.Feed
	json.NewDecoder(r.Body).Decode(&newFeed)

	var errors []string
	if len(newFeed.Name) == 0 {
		errors = append(errors, "The field name required")
	}

	if len(newFeed.Url) == 0 {
		errors = append(errors, "The field url required")
	}

	if len(errors) > 0 {
		response_http.Return400(w, response_http.ErrorMessage{
			StatusCode: 400,
			Error:      errors,
		})
		return
	}

	userId := r.Context().Value("user_id")
	newFeed.UserID = userId
	feedCreated, err := controller.repository.Create(newFeed)
	if err != nil {
		log.Printf("Error => %v", err)
	}

	response_http.ReturnJson(w, 201, feedCreated)
}

func (controller *FeedController) GetFeeds(w http.ResponseWriter, r *http.Request) {
	_, err := controller.repository.GetFeeds()
	if err != nil {
		log.Printf("Error => %v", err)
	}

	feeds, err := controller.repository.GetFeeds()
	if err != nil {
		log.Printf("Error => %v", err)
	}
	response_http.ReturnJson(w, 200, feeds)
}

func (controller *FeedController) Follow(w http.ResponseWriter, r *http.Request) {

	feedId := chi.URLParam(r, "feedId")
	userId := r.Context().Value("user_id")

	feedIdParsed, _ := strconv.ParseInt(feedId, 10, 64)

	newFeed := model.Feed{
		ID:     feedIdParsed,
		UserID: userId,
	}

	_, err := controller.repository.Follow(newFeed)
	if err != nil {
		log.Printf("Error => %v", err)
	}

	response_http.ReturnJson(w, 204, struct{}{})
}

func (controller *FeedController) Unfollow(w http.ResponseWriter, r *http.Request) {

	feedId := chi.URLParam(r, "feedId")
	userId := r.Context().Value("user_id")

	feedIdParsed, _ := strconv.ParseInt(feedId, 10, 64)

	newFeed := model.Feed{
		ID:     feedIdParsed,
		UserID: userId,
	}
	_, err := controller.repository.Unfollow(newFeed)
	if err != nil {
		log.Printf("Error => %v", err)
	}

	response_http.ReturnJson(w, 204, struct{}{})
}
