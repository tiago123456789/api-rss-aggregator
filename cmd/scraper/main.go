package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/config"
	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/repository"
	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/service"
	"github.com/tiago123456789/api-rss-aggregator/pkg/scheduler"
)

const INTERVAL_5MINUTES = 300

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.StartDB()
	if err != nil {
		log.Fatal(err)
	}

	postRepository := repository.NewPostepository(db)
	feedRepository := repository.NewFeedRepository(db)
	rssFeedService := service.NewRssFeedService(postRepository, feedRepository)

	scheduler.Task(
		INTERVAL_5MINUTES,
		rssFeedService.ProcessToUpdateNewPostsRssFeedUrl,
	)
}
