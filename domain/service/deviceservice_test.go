package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDevice(t *testing.T) {
	var expectedAlias = "testDevice"
	var expectedType = "mobilecomputer"
	var expectedModel = "CT60"

	device := CreateDevice(expectedAlias, expectedModel, expectedType)
	assert.Equal(t, expectedModel, device.Model, "should be equal")
	assert.Equal(t, expectedAlias, device.Name, "should be equal")
	assert.Equal(t, expectedType, device.Type, "should be equal")

	assert.NotNil(t, device.SystemID)
}
