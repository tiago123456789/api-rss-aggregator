package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/tiago123456789/api-rss-aggregator/interval/api_rss/repository"
	"github.com/tiago123456789/api-rss-aggregator/pkg/response_http"
)

func HasAuthenticated(
	handler http.HandlerFunc, repository repository.UserRepository,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("authorization")

		if len(apiKey) == 0 {
			message := []string{"UnAuthorized to make request"}
			response_http.ReturnJson(w, 401, response_http.ErrorMessage{
				StatusCode: 401,
				Error:      message,
			})
			return
		}

		apiKey = strings.ReplaceAll(apiKey, "ApiKey ", "")
		user, err := repository.GetByApiKey(apiKey)
		if err != nil {
			fmt.Printf("Error => %v", err)
			message := []string{"You can't make request"}
			response_http.ReturnJson(w, 403, response_http.ErrorMessage{
				StatusCode: 401,
				Error:      message,
			})
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", user.ID)
		newRequest := r.WithContext(ctx)
		handler.ServeHTTP(w, newRequest)
	}
}
