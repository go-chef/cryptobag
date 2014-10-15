package cryptobag

import (
	"fmt"

	"github.com/go-chef/chef"
)

type KeyMap map[string]*EncryptedDataBagItemKey

type EncryptedDataBagItem struct {
	Id   string
	Keys KeyMap
}

type EncryptedDataBagItemKey struct {
	Cipher        string
	EncryptedData string
	Iv            string
	Version       float64
}

func newEncryptedDataBagItemKey(in interface{}) *EncryptedDataBagItemKey {
	return &EncryptedDataBagItemKey{
		Cipher:        in.(map[string]interface{})["cipher"].(string),
		EncryptedData: in.(map[string]interface{})["encrypted_data"].(string),
		Iv:            in.(map[string]interface{})["iv"].(string),
		Version:       in.(map[string]interface{})["version"].(float64),
	}
}

func NewEncryptedDataBagItem(in chef.DataBagItem) (item *EncryptedDataBagItem) {
	item = new(EncryptedDataBagItem)
	item.Keys = make(map[string]*EncryptedDataBagItemKey)

	for k, v := range in.(map[string]interface{}) {
		switch k {
		case "id":
			item.Id = v.(string)
		default:
			item.Keys[k] = newEncryptedDataBagItemKey(v)
		}
	}

	return item
}

func (e *EncryptedDataBagItem) DecryptKey(keyName, secret string) interface{} {
	key := e.Keys[keyName]

	switch key.Version {
	case 1:
		return version1Decoder([]byte(secret), key.Iv, key.EncryptedData)
	default:
		panic(fmt.Sprintf("not implemented for encrypted bag version %d!", key.Version))
	}
}
