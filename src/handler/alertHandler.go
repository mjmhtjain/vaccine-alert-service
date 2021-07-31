package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/mjmhtjain/vaccine-alert-service/src/cowin"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

var cowinService cowin.AppointmentService

func init() {
	cowinService = cowin.NewAppointmentService()
}

func AlertHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		fetchAppointments(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func fetchAppointments(w http.ResponseWriter, r *http.Request) {
	stateName := strings.TrimPrefix(r.URL.Path, "/appointments/")

	log.Printf("stateName: %v", stateName)

	appointments, err := cowinService.FetchVaccineAppointments(stateName, util.TodaysDate())
	if err != nil {
		log.Printf("Error occured while fetching Appointments: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(appointments)
}
