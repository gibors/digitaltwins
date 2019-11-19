package repository

import (
	"log"
	"os"
	"testing"
)

func TestGetInsertedDevice(t *testing.T) {

	cl := MongoClient{}
	AUT_DMSTRING := os.Getenv("DbConnectionString_DM")

	client, err := GetMongoClient(AUT_DMSTRING)
	if err != nil {
		log.Fatalf("Error getting client %v", err)
	}

	cl.ClientAut = client

	cl.DMCollectionAut = client.Database("devicemanagement").Collection("devices")

	device, err := cl.GetDeviceInserted("CT60MAUTXPP7P1")
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
