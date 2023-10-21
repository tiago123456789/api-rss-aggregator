package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/config"
	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/controller"
	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	db, err := config.StartDB()
	if err != nil {
		log.Fatal(err)
	}

	userController := controller.NewUsercontroller(db)
	feedController := controller.NewFeedController(db)
	postController := controller.NewPostRepository(db)
	r.Get("/healthz", userController.Healthz)
	r.Post("/users", userController.Create)
	r.Post("/feeds", middleware.HasAuthenticated(
		feedController.Create, *userController.GetRepository(),
	))
	r.Get("/posts", middleware.HasAuthenticated(
		postController.GetPosts, *userController.GetRepository(),
	))
	r.Get("/feeds", feedController.GetFeeds)
	r.Get("/follow-feeds/{feedId}", middleware.HasAuthenticated(
		feedController.Follow, *userController.GetRepository(),
	))
	r.Get("/unfollow-feeds/{feedId}", middleware.HasAuthenticated(
		feedController.Unfollow, *userController.GetRepository(),
	))

	log.Printf("Server is running at port 8000")
	port := os.Getenv("PORT")

	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}
}
