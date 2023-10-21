package response_http

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorMessage struct {
	StatusCode int      `json:"statusCode"`
	Error      []string `json:"error"`
}

func Return400(w http.ResponseWriter, data ErrorMessage) {
	w.Header().Add("Content-Type", "application/json")
	ReturnJson(w, 400, data)
}

func Return200AndEmpty(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
}

func ReturnJson(w http.ResponseWriter, code int, data interface{}) {
	dataInJson, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to parse data to json response: %v", data)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dataInJson)
}
