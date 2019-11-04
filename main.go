package main

import (
	"caidc_auto_devicetwins/config"
	device "caidc_auto_devicetwins/domain/model"
	"caidc_auto_devicetwins/domain/service"
	"log"
)

func main() {
	tenantID := "5dbe56e6f5b7d751f8940a87"

	log.Println("Device simulation started.. ")

	values := config.GetConfigValues()

	completed, err := service.OnboardDeviceOnPremise(values, "CT60", "MyFirstDevice",
		device.MOBILECOMPUTER, tenantID)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Succesfully onboarded %t", completed)
	}
}
