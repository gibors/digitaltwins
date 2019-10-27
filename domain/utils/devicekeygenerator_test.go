package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeyPair(t *testing.T) {
	keyPair := generateKeyPair()
	assert.NotNil(t, keyPair)
}
