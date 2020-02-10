package service

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecryptString(t *testing.T) {
	encryptedString := "Oy//9UMakumI5JhnyBMMQiWv9VzJjCAun/wiTq2XpLT3ahpyqzdsMCv1Tw6QXRPyoUbxIpX1Orw6g+BkOtlkd3rE51UeEaI4RzLD/ykXSvzUcavMMbOuDPRVuQA4RgeyXWFOotRl5JPsGwNGhLHkuT0LxIGtFFAhwPpMghi1dKk8p4LS0DsEoPtmrbRWbGy5"

	value := DecryptString(encryptedString)
	log.Println("decrypted value: ")
	log.Println(value)
	assert.NotEqual(t, "", value)

}
