package service

import (
	"caidc_auto_devicetwins/config"
	dev "caidc_auto_devicetwins/domain/model"
	"caidc_auto_devicetwins/domain/repository"
	"caidc_auto_devicetwins/domain/utils"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func OnboardDeviceOnPremise(conf config.Configuration, model string,
	dtype string, tenantID string) (bool, error) {
	log.Println("Creating device >>>")

	device := CreateDevice(model, dtype)

	log.Printf("Device created >>> %s", device.SerialNumber)
	log.Println("")

	//TODO: generate signature when creating device
	signature := "MEQCIB+KXWEnlxHAVjJIra/LDM0p3mbPAKM/pRpzcivCx0w/AiAxVUi/aFCDCp0/nB/OKQdtjtIQYcksLbkOPMXhFXqnMQ=="
	deviceID := "CT60_03VR2FVES7"

	repo := repository.Repository{
		ConfigParams: conf,
	}

	log.Println("Getting global token... >>> ")
	repo.GlobalToken = repo.GetGlobalToken(signature, deviceID)

	log.Println("Associating device to tenant..  >>> ")
	repo.AssociteDeviceToAtenant(device.SystemID, tenantID)
	log.Println("Device successfully associated >>> ")

	log.Println("Onboarding device.. >>> ")
	onboardingSucced := repo.OnboardDevice(device)

	var systemGuid = ""

	if onboardingSucced == true {

		log.Printf("Device onbaorded suscessfully %s", device.SerialNumber)
		systemGuid = insertDeviceFromDev(device.SerialNumber)
		infoTenant := repo.GetTenantInformation(device.SystemID)
		log.Println("Tenant information: ")
		log.Println(infoTenant)
		repo.TenantToken = repo.GetTenantToken(signature, deviceID)

		apiEndpoint := FindPropertyByName("api", infoTenant.Properties)
		queueEndpoint := FindPropertyByName("queue", infoTenant.Properties)
		queueKey := FindPropertyByName("queuekey", infoTenant.Properties)

		authAPI := apiEndpoint.Value + "/api/auth"

		queueToken := repo.GetQueueToken(authAPI)
		log.Printf("Queue Toke: %s", queueToken)
		queueConfig := QueueConfig{}
		queueConfig.Key = queueKey.Value
		queueConfig.URL = queueEndpoint.Value
		queueConfig.Token = queueToken
		device.SystemGUID = systemGuid
		queueConfig.PublishEvent("device", GenerateNewConnectionData(device)) // new connection event after onboarding successfully

		return true, nil
	}
	log.Fatalf("Device onbaording failed.. %s", device.SerialNumber)
	return false, fmt.Errorf("Onboarded Failed>>>")
}

func sendEvents() {

}

//TODO: Work around to insert device locally once onboarded on automation( we need to figure this out )
func insertDeviceFromAutomation() {
	cl := repository.MongoClient{}
	DMSTRING := os.Getenv("DbConnectionString_DM")
	DMSTRING = DecryptString(DMSTRING)
	AUTDMSTRING := "mongodb://dmdbadmin:Honeywe!!Up!nThe$ky786@caidcautdb001.westus.cloudapp.azure.com:47017/devicemanagement?readPreference=primary"

	autClient, err := repository.GetMongoClient(AUTDMSTRING)

	utils.FailOnError(err, "Error getting mongo client")

	cl.ClientAut = autClient

	cl.DMCollectionAut = autClient.Database("devicemanagement").Collection("devices")

	device, err := cl.GetDeviceInserted("CT60MAUTXPP7P1")

	utils.FailOnError(err, "Error when getting device")

	client, err := repository.GetMongoClient(DMSTRING)

	utils.FailOnError(err, "Error getting mongo client")

	cl.ClientLocal = client
	cl.DMCollectionLocal = client.Database("devicemanagement").Collection("devices")

	result, err := cl.InsertDevice(device)

	utils.FailOnError(err, "Failed while inserting device to DM db")

	log.Printf("Inserted device successfully %v", result)
}

func insertDeviceFromDev(deviceId string) string {
	cl := repository.MongoClient{}
	DMSTRING := os.Getenv("DbConnectionString_DM")
	DMSTRING = DecryptString(DMSTRING)
	AUTDMSTRING := "mongodb://dmsdkuser:Honeywe!!Up!nThe$ky786@devcaidcdb-01.westus.cloudapp.azure.com:24403/devicemanagement?readPreference=primary"

	autClient, err := repository.GetMongoClient(AUTDMSTRING)

	utils.FailOnError(err, "Error getting mongo client")

	cl.ClientAut = autClient

	cl.DMCollectionAut = autClient.Database("devicemanagement").Collection("devices")

	log.Println("Getting device onboarded from DB... >>> ")

	device, err := cl.GetDeviceInserted(deviceId)

	utils.FailOnError(err, "Error when getting device")

	client, err := repository.GetMongoClient(DMSTRING)

	utils.FailOnError(err, "Error getting mongo client")

	cl.ClientLocal = client
	cl.DMCollectionLocal = client.Database("devicemanagement").Collection("devices")
	ID, _ := primitive.ObjectIDFromHex("5b309711001f0c3fd0ff8dce")
	device.OrganizationUnitKey = &ID
	result, err := cl.InsertDevice(device)

	utils.FailOnError(err, "Failed while inserting device to DM db")

	log.Printf("Inserted device successfully %v", result)
	return device.SystemGUID
}

func FindPropertyByName(name string, props []dev.Property) dev.Property {
	for _, prop := range props {
		if name == prop.Name {
			return prop
		}
	}
	return dev.Property{}
}
