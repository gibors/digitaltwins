package config

import (
	"log"
	"testing"
)

func TestGetConfigFile(t *testing.T) {
	values := GetConfigValues()
	log.Println(values)
}
