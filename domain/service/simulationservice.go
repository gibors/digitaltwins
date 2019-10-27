package service

import (
	"caidc_auto_devicetwins/config"
	"fmt"
)

func OnboardDeviceOnPremise(conf config.Configuration, model string, alias string, dtype string) (bool, error) {
	device := CreateDevice(alias, model, dtype)
	// tokenapi := conf.EndPoints.OnboardingToken
	fmt.Println(device)
	return true, nil
}
