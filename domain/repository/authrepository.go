package repository

import (
	"bytes"
	"caidc_auto_devicetwins/domain/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (r *Repository) GetGlobalToken(signature string, deviceID string) string {

	requestBody, err := json.Marshal(map[string]string{
		"grantType":   "authorization_code",
		"accessToken": signature,
		"clientType":  "DEVICEAPP",
	})
	header := map[string]string{"x-device-serial": deviceID}

	url := r.ConfigParams.EndPoints.AuthAPI.URL
	method := r.ConfigParams.EndPoints.AuthAPI.Method
	tokenEmpty := ""
	request := utils.GenerateRequest(header, url, method, tokenEmpty, requestBody)

	timeout := time.Duration(10 * time.Second)
	client := http.Client{Timeout: timeout}

	resp, err := client.Do(request)

	if err != nil {
		log.Fatalln("Error while getting global token")
		log.Fatalln(err)
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
