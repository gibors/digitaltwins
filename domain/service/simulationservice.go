package service

import (
	"caidc_auto_devicetwins/config"
	"caidc_auto_devicetwins/domain/repository"
	"log"
)

func OnboardDeviceOnPremise(conf config.Configuration, model string,
	alias string, dtype string, tenantID string) (bool, error) {

	device := CreateDevice(alias, model, dtype)

	signature := "MEYCIQDeCdJt8qTLIlnQjSWRHokYb9OetmToD1HkZq48rcITGwIhANFXQzLzCzQp0x5GFH3BjjoM7bLf+tYF49PvN5ZWgCCk"
	deviceID := "CT60_03VR2FVES7"

	repo := repository.Repository{
		ConfigParams: conf,
	}

	globalToken := repo.GetGlobalToken(signature, deviceID)
	repo.GlobalToken = globalToken

	repo.AssociteDeviceToAtenant(device.SystemID, tenantID)

	value := repo.OnboardDevice(device)

	tenantInfo := repo.GetTenantInformation(device.SystemID)
	log.Println(tenantInfo)
	log.Println(value)

	return true, nil
}

func getTenantInfo() {}

func SendEvents() {

}
