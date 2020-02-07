package main

import (
	"caidc_auto_devicetwins/config"
	device "caidc_auto_devicetwins/domain/model"
	"caidc_auto_devicetwins/domain/service"
	"caidc_auto_devicetwins/domain/utils"
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const REQUIRED_PARAMS = "Required parameters Device Createtion simulator: caidc_auto_device_twins {tenantId} {model} {numberOfDevices} /n" +
	"Device events simulation: caidc_auto_device_twins {tenantId} {serialNumber} {queue_event_name} {numberOfEvents}"

func main() {

	var tenantID string

	if len(os.Args) < 2 {
		log.Fatal(REQUIRED_PARAMS)
	}
	log.Println(os.Args)

	tenantID = os.Args[1]

	_, err := primitive.ObjectIDFromHex(tenantID)
	utils.FailOnError(err, "Tenant id format is incorrect")

	simulate := service.SimulationConfig{}
	values := config.GetConfigValues()

	simulate.InitSimulation(values, tenantID) // configuration to start simulation

	log.Println("Configuration file read sucessfully >>> ")

	if len(os.Args) == 4 {

		log.Println(">>> Device creation simulation started.. ")
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

		model := os.Args[2]
		numOfDevices, err := strconv.ParseInt(os.Args[3], 10, 64)
		utils.FailOnError(err, " num of devices should be integer")

		for k := 0; k < int(numOfDevices); k++ {

			completed, err := simulate.OnboardDeviceOnPremise(model, device.MOBILECOMPUTER)

			utils.FailOnError(err, "Failed when creating simulated divice")
			log.Printf("Device created Succesfully:  %t", completed)
		}
		log.Println(" ")
		log.Println(">>> Device creation simulation completed sucessfully >> ")

	} else if len(os.Args) == 5 {

		serialNumber := os.Args[2]
		eventName := os.Args[3]
		numOfEvents, err := strconv.ParseInt(os.Args[4], 10, 64)
		utils.FailOnError(err, "Number of events should be int ")

		log.Println("Calling function to send events.. >>")

		for i := 0; i < int(numOfEvents); i++ {
			simulate.SendEvents(serialNumber, eventName)
		}

	} else {
		log.Fatal(REQUIRED_PARAMS)
	}
}
