package handler

import "net/http"

func Router() *http.ServeMux {
	router := http.NewServeMux()

	// router.Handle("/health", http.HandlerFunc(handler.HealthHandler))
	router.Handle("/appointments/", http.HandlerFunc(AlertHandler))

	return router
}
