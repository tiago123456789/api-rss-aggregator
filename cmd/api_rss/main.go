package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/tiago123456789/api-rss-aggregator/interval/api_rss/config"
	"github.com/tiago123456789/api-rss-aggregator/interval/api_rss/controller"
	"github.com/tiago123456789/api-rss-aggregator/interval/api_rss/middleware"
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

	userController := controller.New(db)

	r.Get("/400", userController.Return400)
	r.Get("/healthz", userController.Healthz)
	r.Get("/hi", middleware.HasAuthenticated(
		userController.SayHi, *userController.GetRepository(),
	))
	r.Post("/users", userController.Create)

	log.Printf("Server is running at port 8000")
	port := os.Getenv("PORT")

	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}
}
