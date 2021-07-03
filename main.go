package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/mjmhtjain/vaccine-alert-service/src/cowin"
	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

func main() {
	appointmentService := cowin.NewAppointmentService()
	stateId := fetchStateId("Delhi")
	todaysDate := util.TodaysDate()

	appointments, err := appointmentService.FetchVaccineAppointments(stateId, todaysDate)
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

func fetchStateId(stateName string) string {
	var data model.States
	fileData, err := util.ReadStaticFile("states.json")
	if err != nil {
		logger.ERROR.Panicf("Error reading file: \n %v \n", stateName)
	}

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		logger.ERROR.Panicf("Error occured while unmarshalling.. \n %v \n", err)
	}

	stateMap := make(map[string]string)
	for _, e := range data.States {

		key := strings.ToLower(e.StateName)
		val := fmt.Sprint(e.StateID)

		stateMap[key] = val
	}

	return stateMap[strings.ToLower(stateName)]
}
