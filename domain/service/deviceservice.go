package service

import (
	"bytes"
	device "caidc_auto_devicetwins/domain/model"
	"math/rand"
	"time"
)

const LENGTH = 6
const CHARSET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const UNDERSCORE = "_"

//Todo: use devicekeygenerator utility
const PUBLIC_KEY = "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEANeSK4az0LE8qOLpss7JF25IROpT2gSc7gtdmONYsEkf1Qe6NJChFmnx4az64WnpprraNUPS3rZZeM5nYxcmkA=="

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
	var serial = "MAUT" + randID
	b.WriteString(model)
	b.WriteString(UNDERSCORE)
	b.WriteString(serial)
	device.SystemID = b.String()
	b.Reset()
	b.WriteString(model)
	b.WriteString(serial)
	device.SerialNumber = b.String()
	device.PublicKey = PUBLIC_KEY

	return device
}

func generateRandomId(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = CHARSET[seededRand.Intn(len(CHARSET))]
	}
	return string(b)
}
