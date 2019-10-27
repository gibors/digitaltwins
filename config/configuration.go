package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	EndPoints Endpoints `json: "Endpoints"`
}

type Endpoints struct {
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
	pwd, _ := os.Getwd()
	var conf Configuration

	jsonFile, err := os.Open(pwd + "/config/config.json")

	if err != nil {
		log.Print("Error reading config file: ")
		log.Fatal(err)
		return conf
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	log.Printf("read file: %s\n", string(byteValue))

	json.Unmarshal(byteValue, &conf)

	return conf
}
