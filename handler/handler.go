package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type SampleResponse struct {
	Time int64 `json:"time"`
}

func TimeHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		response := SampleResponse{
			Time: time.Now().Unix(),
		}
		json.NewEncoder(w).Encode(response)
	}
	return http.HandlerFunc(fn)
}
