package service

import (
	"caidc_auto_devicetwins/config"
	"caidc_auto_devicetwins/domain/repository"
)

func OnboardDeviceOnPremise(conf config.Configuration, model string,
	alias string, dtype string, tenantID string) (bool, error) {

	device := CreateDevice(alias, model, dtype)
	// token := repository.GetOnboardingToken(conf)
	repo := repository.Repository{
		ConfigParams: conf}
	value := repo.OnboardDevice(device)
	if value {
		repo.AssociteDeviceToAtenant(device.SystemID, tenantID)
	}

	return true, nil
}

func SendEvents() {

}
