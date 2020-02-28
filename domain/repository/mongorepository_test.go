package repository

import (
	device "caidc_auto_devicetwins/domain/model"
	"log"
	"testing"
)

func TestGetInsertedDevice(t *testing.T) {

	cl := MongoClient{}
	AUTDMSTRING := "mongodb://dmsdkuser:Honeywe!!Up!nThe$ky786@devcaidcdb-01.westus.cloudapp.azure.com:24403/devicemanagement?readPreference=primary"

	client, err := GetMongoClient(AUTDMSTRING)
	if err != nil {
		log.Fatalf("Error getting client %v", err)
	}

	cl.ClientAut = client

	cl.DMCollectionAut = client.Database("devicemanagement").Collection("devices")

	device, err := cl.GetDeviceInserted("CT60PTR54")
	if err != nil {
		log.Fatal(err)
	}
	// clientLocal, err := getMongoClient(AUT_DMSTRING)

	// cl.ClientLocal = clientLocal
	cl.DMCollectionLocal = cl.DMCollectionAut
	device.SerialNumber = device.SerialNumber + "1"

	result, _ := cl.InsertDevice(device)

	log.Println(result)
}

func TestUpdateDeviceToRegister(t *testing.T) {
	cl := MongoClient{}
	devTR := device.DeviceToRegister{
		DeviceID:            "CT60_PTR55",
		DeviceOwnershipCode: "CT60PTR55",
	}
	device := device.Device{
		SerialNumber: devTR.DeviceOwnershipCode,
		SystemGUID:   "4f1705c0-bb0e-4ebb-8463-b04ff278ab03",
	}
	result := cl.UpdateDeviceToRegister(devTR, device)
	log.Println(result)
}
