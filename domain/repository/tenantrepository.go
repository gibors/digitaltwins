package repository

import (
	"caidc_auto_devicetwins/domain/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (repo *Repository) AssociteDeviceToAtenant(device string, tenantId string) bool {

	url := repo.ConfigParams.EndPoints.AssociateTenant.URL
	method := repo.ConfigParams.EndPoints.AssociateTenant.Method

	req := utils.GenerateRequest(nil, url, method, repo.GlobalToken, nil)

	timeout := time.Duration(10 * time.Second)
	client := http.Client{Timeout: timeout}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("error associating device to tenant status code %d, due to %s",
			resp.StatusCode, resp.Status)
		return false
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
	return true
}

func (rep *Repository) GetTenantInformation(deviceID string) map[string]interface{} {

	url := rep.ConfigParams.EndPoints.GetTenantInfo.URL
	method := rep.ConfigParams.EndPoints.GetTenantInfo.Method

	req := utils.GenerateRequest(nil, url, method, rep.GlobalToken, nil)

	timeout := time.Duration(10 * time.Second)
	client := http.Client{Timeout: timeout}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("error getting tenant information")
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)

	return result
}
