package main

import (
	"caidc_auto_devicetwins/config"
	"caidc_auto_devicetwins/domain/service"
	"log"
)

func main() {

	log.Println("Device simulation started.. ")

	values := config.GetConfigValues()

	completed, err := service.OnboardDeviceOnPremise(values, "CT60", "MyFirstDevice", "mobilecomputer")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Succesfully onboarded %t", completed)
	}
}
