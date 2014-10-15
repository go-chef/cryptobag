package cryptobag

import (
	"encoding/base64"
	"fmt"
)

func decodeBase64(str string) []byte {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
	}
	return data
}

func unPKCS7Padding(data []byte) []byte {
	dataLen := len(data)
	endIndex := int(data[dataLen-1])

	if 16 > endIndex {
		return data[:dataLen-endIndex]
	}
	return nil
}
