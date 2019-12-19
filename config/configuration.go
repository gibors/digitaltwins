package config

import (
	"caidc_auto_devicetwins/domain/utils"
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	EndPoints Endpoints `json: "Endpoints"`
}

type Endpoints struct {
	AuthAPI              Params `json: "AuthApi"`
	GetTenantInfo        Params `json: "GetTenantInfo"`
	AssociateTenant      Params `json: "AssociateTenant"`
	OnboardDeviceMobile  Params `json: "OnboardDeviceMobile"`
	OnboardingToken      Params `json: "OnboardingToken"`
	OnboardDeviceGateway Params `json: "OnboardDeviceGateway"`
}

type Params struct {
	URL    string `json: "url"`
	Method string `json: "method"`
}

func GetConfigValues() Configuration {
	var conf Configuration

	configFile, err := os.Open("./resources/config.json")
	path, _ := os.Getwd()
	log.Printf("path: %s", path)
	utils.FailOnError(err, "Failed when reading config file")

	jsonParser := json.NewDecoder(configFile)

	jsonParser.Decode(&conf)

	return conf
}
