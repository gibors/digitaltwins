package repository

import (
	"bytes"
	"caidc_auto_devicetwins/config"
	device "caidc_auto_devicetwins/domain/model"
	"caidc_auto_devicetwins/domain/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Repository struct {
	ConfigParams config.Configuration
}

func (r *Repository) OnboardDevice(device device.Device) bool {
	tokenDetails := utils.GetTokenDetails()
	tokenReq := make(map[string]interface{})
	tokenReq["QRCodeHash"] = tokenDetails.TokenDetails.TokenHash
	tokenReq["TokenId"] = tokenDetails.TokenDetails.TokenID
	tokenReq["RegistrationDetails"] = map[string]interface{}{
		"Ownershipcode": device.SerialNumber,
		"PublicKey":     device.PublicKey,
		"SystemId":      device.SystemID,
	}
	reqBody, err := json.Marshal(tokenReq)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	req, err := http.NewRequest("POST", r.ConfigParams.EndPoints.OnboardDeviceMobile.URL, bytes.NewBuffer(reqBody))

	req.Header.Set("Authorization", "Bearer "+tokenDetails.TokenDetails.JWTToken)
	req.Header.Set("Content-type", "application/json")
	timeout := time.Duration(100 * time.Second)
	client := http.Client{Timeout: timeout}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
		return false
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	status := result["statusCode"].(string)
	log.Println(result)
	if status != string(16) {
		return false
	}
	return true
}

func GetOnboardingToken(conf config.Configuration) string {

	authToken := GetAuthToken(conf.EndPoints.AuthAPI.URL)

	requestBody, err := json.Marshal(map[string]string{
		"CustomerName": "caidcapiautomationuser@honeywell.com",
		"DeviceCount":  "100",
		"ExpiresAfter": "30",
		"Site":         "5c05b0000000000101100100",
		"SiteName":     "name",
	})

	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", conf.EndPoints.OnboardDeviceMobile.URL,
		bytes.NewBuffer(requestBody))

	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-type", "application/json")
	timeout := time.Duration(10 * time.Second)
	client := http.Client{Timeout: timeout}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)

	return result["responseToken"].(string)
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
