package utils

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGerateEventTimeStamp(t *testing.T) {

	tm := GenerateEventTimeStamp()
	log.Println(tm)
	assert.NotNil(t, tm)
}
