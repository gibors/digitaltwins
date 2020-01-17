package device

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	SerialNumber              string              `bson:"serialNumber" json:"serialNumber"`
	Name                      string              `bson:"name"`
	SystemGUID                string              `bson:"systemGuid" json:"systemGuid"`
	Type                      string              `bson:"type"`
	Model                     string              `bson:"model"`
	ConnectionStatus          string              `bson:"connectionStatus" json:"connectionStatus"`
	Family                    *string             `bson:"family"`
	DisplayModel              *string             `bson:"displayModel" json:"displayModel"`
	PhotoURL                  *string             `bson:"photoURL" json:"photoURL"`
	SoftwareVersion           *string             `bson:"softwareVersion" json:"softwareVersion"`
	Description               *string             `bson:"description"`
	FirmwareVersion           *string             `bson:"firmwareVersion" json:"firmwareVersion"`
	DisplayFirmwareVersion    *string             `bson:"displayFirmwareVersion" json:"displayFirmwareVersion"`
	GatewayID                 *primitive.ObjectID `bson:"gatewayId" json:"gatewayId"`
	OsVersion                 *string             `bson:"osVersion" json:"osVersion"`
	LastOnline                *time.Time          `bson:"lastOnline" json:"lastOnline"`
	LastUpdated               *time.Time          `bson:"lastUpdated" json:"lastUpdated"`
	ProvisionedUserEmail      *string             `bson:"provisionedUserEmail" json:"provisionedUserEmail"`
	DeviceDbStatus            *string             `bson:"deviceDbStatus" json:"deviceDbStatus"`
	Properties                *string             `bson:"properties"`
	BatterySerialNumber       *string             `bson:"batterySerialNumber" json:"batterySerialNumber"`
	GatewayHistory            *[]string           `bson:"gatewayHistory" json:"gatewayHistory"`
	OrganizationUnitKey       *primitive.ObjectID `bson:"organizationUnitKey" json:"organizationUnitKey"`
	SecurityPatchLevel        *string             `bson:"securityPatchLevel" json:"securityPatchLevel"`
	BluetoothAddress          *string             `bson:"bluetoothAddress" json:"bluetoothAddress"`
	WiFiMacAddress            *string             `bson:"wiFiMacAddress" json:"wiFiMacAddress"`
	LanMacAddress             *string             `bson:"lanMacAddress" json:"lanMacAddress"`
	Imei                      *string             `bson:"imei"`
	Meid                      *string             `bson:"meid"`
	CurrentWirelessConnection *string             `bson:"currentWirelessConnection" json:"currentWirelessConnection"`
	RegistrationDatetime      *string             `bson:"registrationDatetime" json:"registrationDatetime"`
	DiscontinuedDatetime      *string             `bson:"discontinuedDatetime" json:"discontinuedDatetime"`
	ProvisioningDatetime      *string             `bson:"provisioningDatetime" json:"provisioningDatetime"`
	Updates                   *[]string           `bson:"updates"`
	Hierarchy                 *string             `bson:"hierarchy"`
	AppsDetails               *[]string           `bson:"appsDetails" json:"appsDetails"`
	CreatredOnDate            time.Time           `bson:"-"`
	PublicKey                 string              `bson:"-"`
	ActiveStatus              bool                `bson:"-"`
	SystemID                  string              `bson:"-"`
	SystemType                string              `bson:"-"`
}

const (
	GATEWAY        = "gateway"
	PRINTER        = "printer"
	SCANNER        = "scanner"
	MOBILECOMPUTER = "mobilecomputer"
)

type KeyPair struct {
	PrivateKey string
	PublicKey  string
}

func (s Device) String() string {
	return fmt.Sprintf("Device [SN: %s - Name: %s - ConnectionStatus:] ", s.SerialNumber, s.Name)
}
