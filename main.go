package main

import (
	"caidc_auto_devicetwins/config"
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

	log.Println("Device simulation started.. ")

	values := config.GetConfigValues()
	log.Println(values)

	// completed, err := service.OnboardDeviceOnPremise(values, "CT60", "MyFirstDevice",
	// 	device.MOBILECOMPUTER, tenantID)

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Printf("Succesfully onboarded %t", completed)
	// }
}
