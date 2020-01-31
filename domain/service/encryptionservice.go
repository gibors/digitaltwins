package service

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func DecryptString(text string) string {
	CIPHER_KEY := []byte("1180228033804480")
	iv := "9980888077806680"
	ciphertext, err := base64.StdEncoding.DecodeString(text)

	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher([]byte(CIPHER_KEY))
	if err != nil {
		panic(err)
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)

	fmt.Printf("%s\n", ciphertext)
	return string(ciphertext)
}
