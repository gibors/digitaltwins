package utils

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTokenDetails(t *testing.T) {

	token := GetTokenDetails()
	log.Println(token)
	assert.NotNil(t, token)
}
