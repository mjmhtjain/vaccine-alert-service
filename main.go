package main

import (
	"fmt"
	"log"

	"github.com/mjmhtjain/vaccine-alert-service/src/cowin"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

func main() {
	appointmentService := cowin.NewAppointmentService()
	appointments, err := appointmentService.FetchVaccineAppointments("141", "25-07-2021")
	if err != nil {
		log.Panicf("Error occured while fetching Appointments: %v\n", err)
	}

	fmt.Println(util.PrettyPrint(appointments))
	// delta := findDelta(appointments)
	// sendNotification(delta)
}

func findDelta(appointments *model.Appointments) string {
	// upserted := upsertAppointments(appointments)
	return ""
}

func sendNotification(appointments string) {
	// fmt.Printf("sent .. %v \n", appointments)
}
