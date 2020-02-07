package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func DecryptString(text string) string {
	SECRET := []byte("#my*S3cr3t")
	iv := "9980888077806680"
	ciphertext, err := base64.StdEncoding.DecodeString(text)
	cipherKey := sha256.Sum256(SECRET)
	key := cipherKey[:]
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher([]byte(key))
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
