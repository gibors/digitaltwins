package service

import (
	"caidc_auto_devicetwins/config"
	dev "caidc_auto_devicetwins/domain/model"
	"caidc_auto_devicetwins/domain/repository"
	"caidc_auto_devicetwins/domain/utils"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SIGNATURE = "MEQCIB+KXWEnlxHAVjJIra/LDM0p3mbPAKM/pRpzcivCx0w/AiAxVUi/aFCDCp0/nB/OKQdtjtIQYcksLbkOPMXhFXqnMQ=="
const DEVICE_ID = "CT60_03VR2FVES7" // DEVICEID
const AUTDMSTRING = "mongodb://dmsdkuser:Honeywe!!Up!nThe$ky786@devcaidcdb-01.westus.cloudapp.azure.com:24403/devicemanagement?readPreference=primary"

type SimulationConfig struct {
	TenantInfo  dev.Tenant
	RepoValue   repository.Repository
	QueueValues QueueConfig
	MongoCl     repository.MongoClient
}

func (s *SimulationConfig) InitSimulation(conf config.Configuration, tenantID string) {
	DMSTRING := os.Getenv("DbConnectionString_DM")
	log.Println(">>> Encrypted DM string : ")
	log.Println(DMSTRING)
	DMSTRING = DecryptString(DMSTRING)
	log.Println(">>> DM local string conn: ")
	log.Println(DMSTRING)
	s.MongoCl = repository.MongoClient{}

	//Get Global Token
	s.RepoValue = repository.Repository{
		ConfigParams: conf,
	}

	s.RepoValue.GlobalToken = s.RepoValue.GetGlobalToken(SIGNATURE, DEVICE_ID)
	s.RepoValue.TenantToken = s.RepoValue.GetTenantToken(SIGNATURE, DEVICE_ID)
	s.TenantInfo = s.RepoValue.GetTenantInformation(tenantID)

	apiEndpoint := FindPropertyByName("SinapsApi", s.TenantInfo.Properties)
	queueEndpoint := FindPropertyByName("queue", s.TenantInfo.Properties)

	authAPI := apiEndpoint.Value + "/api/auth"
	s.QueueValues.Token = s.RepoValue.GetQueueToken(authAPI)
	log.Println(s.QueueValues.Token)
	s.QueueValues.URL = queueEndpoint.Value

	var clientError error
	s.MongoCl.ClientAut, clientError = repository.GetMongoClient(AUTDMSTRING)
	utils.FailOnError(clientError, "Error when creating cloud mongodb client")
	s.MongoCl.DMCollectionAut = s.MongoCl.ClientAut.Database("devicemanagement").Collection("devices")

	s.MongoCl.ClientLocal, clientError = repository.GetMongoClient(DMSTRING)
	utils.FailOnError(clientError, "Error when creating Local mongodb client")
	s.MongoCl.DMCollectionLocal = s.MongoCl.ClientLocal.Database("devicemanagement").Collection("devices")

	log.Println("Successfully simulation initiated  ")
}

func (s *SimulationConfig) OnboardDeviceOnPremise(model string, dtype string) (bool, error) {

	log.Println("Creating device >>>")

	device := CreateDevice(model, dtype)

	log.Printf("Device created >>> %s", device.SerialNumber)
	log.Println("")

	log.Println("Onboarding device.. >>> ")
	onboardingSucced := s.RepoValue.OnboardDevice(device)
	WaitCall()

	var systemGUID = ""

	if onboardingSucced == true {
		log.Printf("Successfully onboarded device: %s ", device.SerialNumber)
		log.Println(" ")

		log.Println("Associating device to tenant..  >>> ")
		s.RepoValue.AssociteDeviceToAtenant(device.SystemID, s.TenantInfo.TenantID.Hex())
		WaitCall()
		log.Println("Device successfully associated to Tenant >>> ")

		systemGUID = s.insertDeviceFromDev(device.SerialNumber)

		log.Printf("systemGUID of inserted device: %s", systemGUID)

		device.SystemGUID = systemGUID
		log.Println("publishing New Connection Event ... >>> ")
		newConnectionEvent := GenerateNewConnectionData(device)
		log.Printf(" Event: %s ", newConnectionEvent)
		log.Println(" ")
		s.QueueValues.PublishEventToRabbit("device", newConnectionEvent) // new connection event after onboarding successfully

		return true, nil
	}
	log.Fatalf("Device onbaording failed.. %s", device.SerialNumber)
	return false, fmt.Errorf("Onboarded Failed>>>")
}

func (s *SimulationConfig) SendEvents(serialNumber string, queue string) bool {
	device, err := s.MongoCl.GetDeviceInserted(serialNumber)
	utils.FailOnError(err, "Error get device from cloud db")

	var message string

	switch queue {
	case dev.TELEMETRYQ:
		message = CreateTelemetryEvent(device)

	default:
		message = CreateTelemetryEvent(device)
	}

	log.Printf("TelemetryMessage: %s", message)
	log.Println("")
	s.QueueValues.PublishEventToRabbit(queue, message)

	return true
}

func (s *SimulationConfig) insertDeviceFromDev(deviceId string) string {

	log.Println("Getting device onboarded from DB... >>> ")

	device, err := s.MongoCl.GetDeviceInserted(deviceId)

	utils.FailOnError(err, "Error when getting device")

	ID, _ := primitive.ObjectIDFromHex("5b309711001f0c3fd0ff8dce")
	device.OrganizationUnitKey = &ID
	result, err := s.MongoCl.InsertDevice(device)

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

func WaitCall() {
	time.Sleep(1 * time.Second)
}
