package service

import (
	"caidc_auto_devicetwins/config"
	"caidc_auto_devicetwins/domain/repository"
	"log"
)

func OnboardDeviceOnPremise(conf config.Configuration, model string,
	alias string, dtype string, tenantID string) (bool, error) {

	device := CreateDevice(alias, model, dtype)

	signature := "MEQCIAb6DNjdsxDh+flzqXQpYIB5gQvwMEjcURmbB3lRIjzhAiAw7L+nOf2qnrR7+gZQR8HGx5o3qLKYFsqyTtN54TxKbw=="
	deviceID := "CN80_17340D8241"

	repo := repository.Repository{
		ConfigParams: conf,
	}

	globalToken := repo.GetGlobalToken(signature, deviceID)
	repo.GlobalToken = globalToken

	repo.AssociteDeviceToAtenant(device.SystemID, tenantID)

	value := repo.OnboardDevice(device)

	tenantInfo = repo.GetTenantInformation(device.SystemID)

	log.Println(value)

	return true, nil
}

func getTenantInfo() {}

func SendEvents() {

}
