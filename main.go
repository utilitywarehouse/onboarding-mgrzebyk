package main

import (
	"net/http"
	"onboarding-mgrzebyk/handler"
)

func main() {
	mux := http.NewServeMux()
	th := handler.TimeHandler()
	mux.Handle("/time", th)

	_ = http.ListenAndServe("0.0.0.0:8080", mux)
}
