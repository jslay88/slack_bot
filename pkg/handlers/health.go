package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Health Check.")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Healthy")
}
