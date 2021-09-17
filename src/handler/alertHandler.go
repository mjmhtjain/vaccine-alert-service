package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/mjmhtjain/vaccine-alert-service/src/cowin"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

type AppointmentController interface {
	AppoitnmentAlertHandler(w http.ResponseWriter, r *http.Request)
}

type AppointmentControllerImpl struct {
	cowinService cowin.AppointmentService
}

func NewAppointmentController() AppointmentController {
	return &AppointmentControllerImpl{
		cowinService: cowin.NewAppointmentService(),
	}
}

func (ctrl *AppointmentControllerImpl) AppoitnmentAlertHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		ctrl.fetchAppointments(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func (ctrl *AppointmentControllerImpl) fetchAppointments(w http.ResponseWriter, r *http.Request) {
	stateName := strings.TrimPrefix(r.URL.Path, "/appointments/")

	log.Printf("stateName: %v", stateName)

	ctx, cancel := context.WithTimeout(r.Context(), util.SLATimeout)
	defer cancel()

	appointments, err := ctrl.cowinService.FetchVaccineAppointments(ctx, stateName, util.TodaysDate())
	if err != nil {
		log.Printf("Error occured while fetching Appointments: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(appointments)
}
