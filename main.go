package main

import (
	"caidc_auto_devicetwins/config"
	device "caidc_auto_devicetwins/domain/model"
	"caidc_auto_devicetwins/domain/service"
	"log"
	"os"
)

func main() {
	var tenantID string
	if len(os.Args) > 1 {
		tenantID = os.Args[1]
		log.Println(tenantID)
	} else {
		log.Fatal("missing TenantID, it should be passed as parameter")
	}

	log.Println(">>> Device simulation started.. ")
	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	values := config.GetConfigValues()
	log.Println("Configuration file read >>> ")

	completed, err := service.OnboardDeviceOnPremise(values, "CT60",
		device.MOBILECOMPUTER, tenantID)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(">>>>")
		log.Printf("Device simulation Succesfully:  %t", completed)
		log.Println(">>>>")
	}
}
