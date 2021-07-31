package handler

import "net/http"

func Router() *http.ServeMux {
	router := http.NewServeMux()
	appController := NewAppointmentController()

	// TODO: router.Handle("/health", http.HandlerFunc(handler.HealthHandler))
	router.Handle("/appointments/", http.HandlerFunc(appController.AppoitnmentAlertHandler))

	return router
}
