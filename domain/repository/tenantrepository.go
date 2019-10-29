package repository

import (
	"bytes"
	device "caidc_auto_devicetwins/domain/model"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (repo *Repository) AssociteDeviceToAtenant(device string, tenantId string) bool {

	req, err := http.NewRequest("POST", repo.ConfigParams.EndPoints.AssociateTenant.URL,
		bytes.NewBuffer(requestBody))

	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-type", "application/json")
	timeout := time.Duration(10 * time.Second)
	client := http.Client{Timeout: timeout}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return true
}

func GetTenantInformation(deviceID string) device.Tenant {
	tenant := device.Tenant{}
	return tenant
}
