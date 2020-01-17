package utils

import (
	"encoding/json"
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
	var tokenConf OnboardingToken

	jsonFile, err := os.Open("./resources/token.json")

	FailOnError(err, "failed to read file")

	jsonParser := json.NewDecoder(jsonFile)

	jsonParser.Decode(&tokenConf)

	return tokenConf

}
