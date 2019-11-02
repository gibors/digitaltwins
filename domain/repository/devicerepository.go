package repository

import (
	"bytes"
	"caidc_auto_devicetwins/config"
	device "caidc_auto_devicetwins/domain/model"
	"caidc_auto_devicetwins/domain/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Repository struct {
	ConfigParams config.Configuration
	GlobalToken  string
}

func (r *Repository) OnboardDevice(dev device.Device) bool {

	tokenDetails := utils.GetTokenDetails()

	requestBody := map[string]interface{}{
		"QRCodeHash": tokenDetails.TokenDetails.TokenHash,
		"TokenId":    tokenDetails.TokenDetails.TokenID,
		"RegistrationDetails": map[string]string{
			"Ownershipcode": dev.SerialNumber,
			"PublicKey":     dev.PublicKey,
			"SystemId":      dev.SystemID,
		},
	}
	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal("Error on creating request body: ")
		log.Fatal(err.Error())
	}
	var url string
	method := r.ConfigParams.EndPoints.OnboardDeviceMobile.Method
	if dev.Type == device.GATEWAY {
		url = r.ConfigParams.EndPoints.OnboardDeviceGateway.URL
	} else if dev.Type == device.MOBILECOMPUTER {
		url = r.ConfigParams.EndPoints.OnboardDeviceMobile.URL
	}

	req := utils.GenerateRequest(nil, url, method, tokenDetails.TokenDetails.JWTToken, reqBody)

	timeout := time.Duration(100 * time.Second)
	client := http.Client{Timeout: timeout}

	log.Println("Request header : ")
	log.Println(req.Header)
	log.Println("Request body : ")
	log.Println(req.Body)

	response, err := client.Do(req)

	if err != nil {
		log.Fatalln("Error while onboarding device: ")
		log.Fatalln(err)
		return false
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatal("Error: ")
		log.Fatal(response.StatusCode)
		log.Fatal(response.Status)
		return false
	}

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)

	status := result["statusCode"].(int)

	log.Printf("result: %s", result)

	if status != 16 {
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
