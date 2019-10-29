package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type OnboardingToken struct {
	TokenDetails TokenDetail `json: "TokenDetails"`
}

type TokenDetail struct {
	TokenID   string `json: "TokenId"`
	JWTToken  string `json: "JWTToken"`
	TokenHash string `json: "TokenHash"`
}

func GetTokenDetails() OnboardingToken {
	pwd, _ := os.Getwd()
	var tokenConf OnboardingToken

	jsonFile, err := os.Open(pwd + "/domain/utils/token.json")

	if err != nil {
		log.Print("Error reading config file: ")
		log.Fatal(err)
		return tokenConf
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &tokenConf)

	return tokenConf
}
