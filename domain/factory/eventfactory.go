package factory

import (
	e "caidc_auto_devicetwins/domain/model"
	u "caidc_auto_devicetwins/domain/utils"
)

func CreateNewConnectionEvent(dev e.Device) e.EventMessage {
	newConnEvent := e.EventMessage{}
	pe := e.CloudPlatformEvent{}
	pe.CreatedTime = u.GenerateEventTimeStamp()
	pe.ID = nil
	pe.CreatorID = nil
	pe.GeneratorID = nil
	pe.GeneratorType = nil
	targeID := "16:79:50:B2:6F:86"
	pe.TargetID = &targeID
	pe.TargetType = nil
	pe.TargetContext = nil
	pe.BodyMessage = nil
	eventType := "NewConnectionEvent"
	pe.EventType = &eventType
	pe.BodyProperties = generateBodyProperties(dev)
	newConnEvent.PlatformEventMessage = pe
	newConnEvent.AnnotationStreamIds = ""

	return newConnEvent
}

func generateBodyProperties(dev e.Device) []e.BodyProperty {
	category := "Device"
	value := "10.77.42.14"
	firmwareVersion := "Ver:19.07_HM2_0160 0006"
	wifimacaddress := "84:25:3F:2E:86:9E"
	return []e.BodyProperty{
		e.BodyProperty{
			Value: &category,
			Key:   "Category",
		},
		e.BodyProperty{
			Key:   "DeviceUniqueID",
			Value: &dev.SerialNumber,
		},
		e.BodyProperty{
			Key:   "TimeStamp",
			Value: nil,
		},
		e.BodyProperty{
			Key:   "Type",
			Value: &dev.Type,
		},
		e.BodyProperty{
			Key:   "Family",
			Value: dev.Family,
		},
		e.BodyProperty{
			Key:   "Name",
			Value: &dev.Name,
		},
		e.BodyProperty{
			Key:   "Model",
			Value: &dev.Model,
		},
		e.BodyProperty{
			Key:   "gatewayconnectivity",
			Value: IsGatewayConnectivity(dev.Type),
		},
		e.BodyProperty{
			Key:   "IPAddress",
			Value: &value,
		},
		e.BodyProperty{
			Key:   "FirmwareVersion",
			Value: &firmwareVersion,
		},
		e.BodyProperty{
			Key:   "OSInfo",
			Value: nil,
		},
		e.BodyProperty{
			Key:   "wifimacaddress",
			Value: &wifimacaddress,
		},
		e.BodyProperty{
			Key:   "Useremail",
			Value: dev.ProvisionedUserEmail,
		},
		e.BodyProperty{
			Key:   "SystemType",
			Value: &dev.SystemType,
		},
		e.BodyProperty{
			Key:   "SystemGuid",
			Value: &dev.SystemGUID,
		},
	}
}

func IsGatewayConnectivity(dtype string) *string {
	value := "true"
	if dtype == e.MOBILECOMPUTER {
		value = "false"
	}
	return &value
}
