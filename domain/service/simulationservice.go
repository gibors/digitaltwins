package service

import (
	"caidc_auto_devicetwins/config"
	dev "caidc_auto_devicetwins/domain/model"
	"caidc_auto_devicetwins/domain/repository"
	"caidc_auto_devicetwins/domain/utils"
	"log"
	"os"
)

func OnboardDeviceOnPremise(conf config.Configuration, model string,
	alias string, dtype string, tenantID string) (bool, error) {

	// device := CreateDevice(alias, model, dtype)

	//TODO: generate signature when creating device
	signature := "MEQCIAb6DNjdsxDh+flzqXQpYIB5gQvwMEjcURmbB3lRIjzhAiAw7L+nOf2qnrR7+gZQR8HGx5o3qLKYFsqyTtN54TxKbw=="
	deviceID := "CN80_17340D8241"
	NEWCONNECTIONMESSAGE := "newconnectionmessage"

	repo := repository.Repository{
		ConfigParams: conf,
	}

	repo.GlobalToken = repo.GetGlobalToken(signature, deviceID)

	// repo.AssociteDeviceToAtenant(device.SystemID, tenantID)

	// value := repo.OnboardDevice(device)

	// if value == true {
	// 	log.Printf("Device onbaorded suscessfully.. %s", device.SerialNumber)
	// 	insertDeviceFromAutomation()
	// } else {
	// 	log.Fatalf("Device onbaording failed.. %s", device.SerialNumber)
	// }

	infoTenant := repo.GetTenantInformation(deviceID)

	repo.TenantToken = repo.GetTenantToken(signature, deviceID)

	apiEndpoint := FindPropertyByName("api", infoTenant.Properties)
	queueEndpoint := FindPropertyByName("queue", infoTenant.Properties)
	queueKey := FindPropertyByName("queuekey", infoTenant.Properties)

	authAPI := apiEndpoint.Value + "/api/auth"

	queueToken := repo.GetQueueToken(authAPI)
	queueConfig := QueueConfig{}
	queueConfig.Key = queueKey.Value
	queueConfig.URL = queueEndpoint.Value
	queueConfig.Token = queueToken
	queueConfig.PublishEvent("device", GetMessageEvent(dev.Device{}, NEWCONNECTIONMESSAGE)) // new connection event after onboarding successfully

	return true, nil
}

//TODO: Work around to insert device locally once onboarded on automation( we need to figure this out )
func insertDeviceFromAutomation() {
	cl := repository.MongoClient{}
	DMSTRING := os.Getenv("DbConnectionString_DM")
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

func FindPropertyByName(name string, props []dev.Property) dev.Property {
	for _, prop := range props {
		if name == prop.Name {
			return prop
		}
	}
	return dev.Property{}
}
