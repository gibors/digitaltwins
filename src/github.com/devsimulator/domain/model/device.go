package device

import (
	"fmt"
	"time"
)

type Device struct {
	Name                 string
	SerialNumber         string
	SystemGUID           string
	Type                 DeviceType
	Model                string
	ConnectionStatus     string
	Family               string
	DisplayModel         string
	PhotoUrl             string
	SoftwareVersion      string
	OsVersion            string
	LastOnline           time.Time
	LastUpdated          time.Time
	CreatredOnDate       time.Time
	ProvisionedUserEmail string
	ActiveStatus         bool
}

type DeviceType string

const (
	GATEWAY        = "gateway"
	PRINTER        = "printer"
	SCANNER        = "scanner"
	MOBILECOMPUTER = "mobilecomputer"
)

func (s Device) String() string {
	return fmt.Sprintf("Device [SN: %s - Name: %s - ConnectionStatus:] ", s.SerialNumber, s.Name)
}
