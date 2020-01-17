package factory

import (
	device "caidc_auto_devicetwins/domain/model"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewConnectionEvent(t *testing.T) {
	dev := device.Device{}
	event := CreateNewConnectionEvent(dev)
	log.Println(event)
	assert.NotNil(t, event.PlatformEventMessage)

}
