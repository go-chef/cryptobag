package cryptobag

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type version1Item struct {
	Content interface{} `json:"json_wrapper"`
}

func version1Decoder(key []byte, iv, encryptedData string) interface{} {
	ciphertext := decodeBase64(encryptedData)
	initVector := decodeBase64(iv)
	keySha := sha256.Sum256(key)

	block, err := aes.NewCipher(keySha[:])
	if err != nil {
		fmt.Println(err)
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, initVector)
	mode.CryptBlocks(ciphertext, ciphertext)

	ciphertext = unPKCS7Padding(ciphertext)

	var item version1Item
	_ = json.Unmarshal(ciphertext, &item)

	return item.Content
}
