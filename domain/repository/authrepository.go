package repository

import (
	"bytes"
	"caidc_auto_devicetwins/domain/utils"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const GDEVICEAPP = "GDEVICEAPP"
const DEVICEAPP = "DEVICEAPP"
const BROKERAPP = "BROKER-RABBITMQ"

func (r *Repository) GetGlobalToken(signature string, deviceID string) string {
	return r.GetToken(GDEVICEAPP, signature, deviceID, nil)
}

func (r *Repository) GetTenantToken(signature string, deviceID string) string {
	return r.GetToken(DEVICEAPP, signature, deviceID, nil)
}

func (r *Repository) GetQueueToken(queueEndpoint string) string {
	empty := ""
	var api *string
	api = &queueEndpoint
	return r.GetToken(BROKERAPP, r.TenantToken, empty, api)
}

func (r *Repository) GetToken(clientType string, accessToken string, deviceID string, urlAuth *string) string {

	requestBody, err := json.Marshal(map[string]string{
		"grantType":   "authorization_code",
		"accessToken": accessToken,
		"clientType":  clientType,
	})
	header := map[string]string{"x-device-serial": deviceID}
	var url string
	if urlAuth == nil {
		url = r.ConfigParams.EndPoints.AuthAPI.URL
	} else {
		url = *urlAuth
		header = nil // for rabbit we should not send headers
	}
	method := r.ConfigParams.EndPoints.AuthAPI.Method
	tokenEmpty := ""
	request := utils.GenerateRequest(header, url, method, tokenEmpty, requestBody)

	timeout := time.Duration(100 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Timeout: timeout, Transport: tr}

	resp, err := client.Do(request)

	if err != nil {
		log.Fatalf("Error while getting global token: %v", err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	return result["accessToken"].(string)
}

func GetAuthToken(url string) string {

	requestBody, err := json.Marshal(map[string]string{
		"grantType":   "authorization_code",
		"accessToken": "vRAEiMGp0jmrcxpv",
		"clientType":  "APIAPP",
	})

	if err != nil {
		log.Fatalln(err)
	}

	timeout := time.Duration(10 * time.Second)
	client := http.Client{Timeout: timeout}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("apikey", "IUcT24epDTAEuhYbqRE1O4l1GIKw8rC6")
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)

	return result["accessToken"].(string)
}
