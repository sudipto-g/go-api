package main

import (
	"fmt"
	"go-api/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting GO API Service")
	err := http.ListenAndServe("localhost:4000", r)
	if err != nil {
		log.Error(err)
	}
}