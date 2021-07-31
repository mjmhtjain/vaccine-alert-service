package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/mjmhtjain/vaccine-alert-service/src/handler"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

const (
	DEFAULT_ADDR string = ":80"
)

//go:embed staticData/*
var embededFiles embed.FS

func init() {
	util.EmbededFiles = embededFiles
}

func main() {
	r := handler.Router()
	log.Fatal(http.ListenAndServe(DEFAULT_ADDR, r))

	// appointmentService := cowin.NewAppointmentService()
	// appointments, err := appointmentService.FetchVaccineAppointments("Delhi", util.TodaysDate())
	// if err != nil {
	// 	log.Panicf("Error occured while fetching Appointments: %v\n", err)
	// }

	// fmt.Println(util.PrettyPrint(appointments))
	// // delta := findDelta(appointments)
	// // sendNotification(delta)
}

func findDelta(appointments *model.Appointments) string {
	// upserted := upsertAppointments(appointments)
	return ""
}

func sendNotification(appointments string) {
	// fmt.Printf("sent .. %v \n", appointments)
}
