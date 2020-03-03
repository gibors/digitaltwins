package service

import (
	"caidc_auto_devicetwins/config"
	dev "caidc_auto_devicetwins/domain/model"
	device "caidc_auto_devicetwins/domain/model"
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
// const AUTDMSTRING = "mongodb://gibadmin:sirenito88@localhost:27017/devicemanagement?readPreference=primary"

const AUTDMSTRING = "mongodb://dmsdkuser:Honeywe!!Up!nThe$ky786@devcaidcdb-01.westus.cloudapp.azure.com:24403/devicemanagement?readPreference=primary"

type SimulationConfig struct {
	TenantInfo dev.Tenant
	RepoValue  repository.Repository
	QueueConf  QueueConfig
	MqttConf   MQTTConfig
	MongoCl    repository.MongoClient
}

func (s *SimulationConfig) InitSimulation(conf config.Configuration, tenantID string) {

	//Get Global Token
	s.RepoValue = repository.Repository{
		ConfigParams: conf,
	}
	var dmConnectionString string
	var clientError error

	if tenantID != "" {
		s.RepoValue.GlobalToken = s.RepoValue.GetGlobalToken(SIGNATURE, DEVICE_ID)
		s.RepoValue.TenantToken = s.RepoValue.GetTenantToken(SIGNATURE, DEVICE_ID)
		s.TenantInfo = s.RepoValue.GetTenantInformation(tenantID)

		apiEndpoint := FindPropertyByName("SinapsApi", s.TenantInfo.Properties)
		queueEndpoint := FindPropertyByName("queue", s.TenantInfo.Properties)

		if s.TenantInfo.Endpoint == "MQTT-RabbitMQ" {
			DMSTRING := os.Getenv("DbConnectionString_DM")
			log.Println(">>> Encrypted DM string : ")
			log.Println(DMSTRING)
			DMSTRING = DecryptString(dmConnectionString)
			log.Println(">>> DM local string conn: ")
			log.Println(DMSTRING)
			s.MongoCl = repository.MongoClient{}
			s.MongoCl.ClientLocal, clientError = repository.GetMongoClient(dmConnectionString)
			utils.FailOnError(clientError, "Error when creating Local mongodb client")
			s.MongoCl.DMCollectionLocal = s.MongoCl.ClientLocal.Database("devicemanagement").Collection("devices")

			authAPI := apiEndpoint.Value + "/api/auth"
			s.QueueConf.Token = s.RepoValue.GetQueueToken(authAPI)
			s.QueueConf.URL = queueEndpoint.Value
			s.RepoValue.ConfigParams.EndPoints.OnboardDeviceGateway.URL = apiEndpoint.Value + "/api/deviceonboarding/cloudgateway"
			s.RepoValue.ConfigParams.EndPoints.OnboardDeviceMobile.URL = apiEndpoint.Value + "/api/deviceonboarding/mobilecomputers"

			log.Println(s.QueueConf.Token)
		} else if s.TenantInfo.Endpoint == "IotHub" {
			s.MqttConf.connectionString = s.RepoValue.GetIotHubToken(DEVICE_ID)
			s.MqttConf.endpoint = queueEndpoint.Value
		}
	}
	s.MongoCl.ClientAut, clientError = repository.GetMongoClient(s.RepoValue.ConfigParams.EndPoints.DbConnectionString.URL)
	utils.FailOnError(clientError, "Error when creating cloud mongodb client")
	s.MongoCl.DMCollectionAut = s.MongoCl.ClientAut.Database("devicemanagement").Collection("devices")

	log.Println("Successfully simulation initiated  ")
}

func (s *SimulationConfig) OnboardDeviceOnPremise(model string, dtype string) (bool, error) {

	var systemGUID = ""
	device := s.OnboarDeviceToForge(model, dtype)

	if device.SystemGUID != "" {
		log.Printf("Successfully onboarded device: %s ", device.SerialNumber)
		log.Println(" ")

		log.Println("Associating device to tenant..  >>> ")
		s.RepoValue.AssociteDeviceToAtenant(device.SystemID, s.TenantInfo.TenantID.Hex())
		WaitCall()
		log.Println("Device successfully associated to Tenant >>> ")

		devInserted, err := s.MongoCl.GetDeviceInserteOnPrem(device.SerialNumber)
		utils.FailOnError(err, "Failed getting device from onprem DB")
		systemGUID = devInserted.SystemGUID
		log.Printf("systemGUID of inserted device: %s", systemGUID)

		device.SystemGUID = systemGUID
		log.Println("publishing New Connection Event ... >>> ")
		newConnectionEvent := GenerateNewConnectionData(device)
		log.Printf(" Event: %s ", newConnectionEvent)
		log.Println(" ")
		s.QueueConf.PublishEventToRabbit("device", newConnectionEvent) // new connection event after onboarding successfully
		return true, nil
	}
	log.Fatalf("Device onbaording failed.. %s", device.SerialNumber)
	return false, fmt.Errorf("Onboarded Failed >>>")
}

func (s *SimulationConfig) OnboarDeviceToForge(model string, dtype string) device.Device {
	log.Println("Creating device >>>")
	var deviceToRegister device.DeviceToRegister
	var deviceInserted device.Device
	var device device.Device

	if s.TenantInfo.TenantName == "" {
		deviceToRegister = s.MongoCl.GetDevicesFromDeviceToRegisterCollection()
		device = CreateDeviceFromDeviceToRegister(deviceToRegister, model, dtype)
	}
	log.Printf("Device created >>> %s", device.SerialNumber)
	log.Println("")

	log.Println("Onboarding device.. >>> ")
	onboardingSucced := s.RepoValue.OnboardDevice(device)

	WaitCall()
	if onboardingSucced && s.TenantInfo.Endpoint != "MQTT-RabbitMQ" {
		var err error
		log.Println("Getting device onboarded from DB... >>> ")

		deviceInserted, err = s.MongoCl.GetDeviceInserted(deviceToRegister.DeviceOwnershipCode)
		if s.TenantInfo.TenantName == "" {
			s.MongoCl.UpdateDeviceToRegister(deviceToRegister, deviceInserted)
		}
		utils.FailOnError(err, "Error when getting device")
		return deviceInserted
	}
	return deviceInserted
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

	if s.TenantInfo.Endpoint == "IotHub" {
		s.MqttConf.PublishMessage(message, device.SystemGUID)
	} else {
		s.QueueConf.PublishEventToRabbit(queue, message)
	}

	return true
}

func (s *SimulationConfig) insertDeviceFromDev(device device.Device) string {

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
