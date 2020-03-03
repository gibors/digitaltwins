package config

import (
	"caidc_auto_devicetwins/domain/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	EndPoints    Endpoints   `json: "Endpoints"`
	TokenDetails TokenDetail `json: "TokenDetails"`
	Environment  string
}
type TokenDetail struct {
	TokenID   string `json: "TokenId"`
	JWTToken  string `json: "JWTToken"`
	TokenHash string `json: "TokenHash"`
}

type Endpoints struct {
	AuthAPI              Params `json: "AuthApi"`
	GetTenantInfo        Params `json: "GetTenantInfo"`
	AssociateTenant      Params `json: "AssociateTenant"`
	OnboardDeviceMobile  Params `json: "OnboardDeviceMobile"`
	OnboardingToken      Params `json: "OnboardingToken"`
	OnboardDeviceGateway Params `json: "OnboardDeviceGateway"`
	DbConnectionString   Params `json: "DbConnectionString"`
}

type Params struct {
	URL    string `json: "url"`
	Method string `json: "method"`
}

const (
	DEV   = "dev"
	QA    = "qa"
	AUT   = "aut"
	LOCAL = "local"
)

func GetConfigValues(env string) Configuration {
	var conf Configuration
	var envName string

	switch env {
	case LOCAL:
		envName = "config-local"
	case DEV:
		envName = "config-dev"
	case QA:
		envName = "config-qa"
	case AUT:
		envName = "config-aut"
	default:
		envName = "config-dev"
	}

	log.Printf("selected environment.. %s ", envName)

	configFile, err := os.Open(fmt.Sprintf("./resources/%s.json", envName))

	path, _ := os.Getwd()
	log.Printf("path: %s", path)
	utils.FailOnError(err, "Failed when reading config file")

	jsonParser := json.NewDecoder(configFile)

	jsonParser.Decode(&conf)
	conf.Environment = env

	return conf
}
