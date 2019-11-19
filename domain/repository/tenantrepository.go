package repository

import (
	dev "caidc_auto_devicetwins/domain/model"
	"caidc_auto_devicetwins/domain/utils"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

func (repo *Repository) AssociteDeviceToAtenant(device string, tenantId string) bool {

	url := repo.ConfigParams.EndPoints.AssociateTenant.URL
	method := repo.ConfigParams.EndPoints.AssociateTenant.Method
	url = strings.Replace(url, "{tenantID}", tenantId, 1)
	url = strings.Replace(url, "{deviceID}", device, 1)
	req := utils.GenerateRequest(nil, url, method, repo.GlobalToken, nil)

	timeout := time.Duration(50 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Timeout: timeout, Transport: tr}
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

func (rep *Repository) GetTenantInformation(deviceID string) dev.Tenant {

	url := rep.ConfigParams.EndPoints.GetTenantInfo.URL
	method := rep.ConfigParams.EndPoints.GetTenantInfo.Method
	url = strings.Replace(url, "{deviceID}", deviceID, 1)
	req := utils.GenerateRequest(nil, url, method, rep.GlobalToken, nil)

	timeout := time.Duration(10 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Timeout: timeout, Transport: tr}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("error getting tenant information")
	}

	var result = dev.Tenant{}
	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)

	return result
}
