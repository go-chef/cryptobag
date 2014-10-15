package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chef/chef"
	"github.com/go-chef/cryptobag"
)

func main() {
	clientName := "bigkraig"
	baseUrl := "https://chefsrever"
	bagName := "testbag"
	itemName := "testitem"
	itemKeyName := "other"
	secret := "testsecret"

	// read a client key
	key, err := ioutil.ReadFile("key.pem")
	if err != nil {
		fmt.Println("Couldn't read key.pem:", err)
		os.Exit(1)
	}

	// build a client
	client, err := chef.NewClient(&chef.Config{
		Name:    clientName,
		Key:     string(key),
		BaseURL: baseUrl,
		SkipSSL: true,
	})
	if err != nil {
		fmt.Println("Issue setting up client:", err)
		os.Exit(1)
	}

	item, err := client.DataBags.GetItem(bagName, itemName)
	if err != nil {
		log.Fatal("Couldn't get item: ", err)
	}

	encrypteditem := cryptobag.NewEncryptedDataBagItem(item)

	f := encrypteditem.DecryptKey(itemKeyName, secret)

	spew.Dump(f)
}
