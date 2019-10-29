package main

import (
	"caidc_auto_devicetwins/config"
	"caidc_auto_devicetwins/domain/service"
	"log"
)

func main() {
	tenantID := "tenantid"

	log.Println("Device simulation started.. ")

	values := config.GetConfigValues()

	completed, err := service.OnboardDeviceOnPremise(values, "CT60", "MyFirstDevice",
		"mobilecomputer", tenantID)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Succesfully onboarded %t", completed)
	}
}
