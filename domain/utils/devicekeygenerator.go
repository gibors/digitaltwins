package utils

import (
	device "caidc_auto_devicetwins/domain/model"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"log"
)

func generateKeyPair() device.KeyPair {
	keys := device.KeyPair{}
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	x509EncodedPV, _ := x509.MarshalECPrivateKey(privateKey)
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(privateKey.PublicKey)

	pvKey := base64.StdEncoding.EncodeToString(x509EncodedPV)
	puKey := base64.StdEncoding.EncodeToString(x509EncodedPub)

	log.Println(pvKey)
	log.Println(puKey)

	return keys
}
