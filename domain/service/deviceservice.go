package service

import (
	"bytes"
	device "caidc_auto_devicetwins/domain/model"
	"math/rand"
	"time"
)

const LENGTH = 10
const CHARSET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const UNDERSCORE = "_"

// const MOBILECOMPUTER = "mobilecomputer"
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func CreateDevice(alias string, model string, deviceType string) device.Device {

	randID := generateRandomId(LENGTH)
	var device = device.Device{
		Model: model,
		Type:  deviceType,
		Name:  alias}

	var b bytes.Buffer
	b.WriteString(model)
	b.WriteString(UNDERSCORE)
	b.WriteString(randID)
	device.SystemID = b.String()
	b.Reset()
	b.WriteString(model)
	b.WriteString(randID)
	device.SerialNumber = b.String()

	return device
}

func generateRandomId(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = CHARSET[seededRand.Intn(len(CHARSET))]
	}
	return string(b)
}
