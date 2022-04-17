package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var DC DonatorConfig

func loadConfig() {
	f, err := os.ReadFile("./donator.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(f, &DC)
	if err != nil {
		panic(err)
	}
}

func writeConfig() {
	d, err := json.Marshal(DC)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("./donator.json", d, 0644)
	if err != nil {
		panic(err)
	}
}

func init() {
	loadConfig()
}

func GenerateDonatorKeyStub(encKey []byte) (string, error) {
	k := keyStub{
		Data: struct {
			Id int `json:"id"`
		}{
			Id: DC.KeyStubID,
		},
		Time: time.Now().Unix(),
	}
	DC.KeyStubID++
	DC.Stubs = append(DC.Stubs, k)
	writeConfig()
	key, err := json.Marshal(k)
	if err != nil {
		return "", err
	}
	enc, err := Encrypt(key, encKey)
	if err != nil {
		return "", err
	}
	return enc, nil
}

func RedeemKeyStub(ks, ip, version string, encKey, keyGenKey []byte) string {
	loadConfig()
	dec, err := Decrypt(ks, encKey)
	if err != nil {
		return ""
	}
	var k keyStub
	err = json.Unmarshal([]byte(dec), &k)
	if err != nil {
		return ""
	}
	var inArray = false
	for i, v := range DC.Stubs {
		if v.Data.Id == k.Data.Id {
			inArray = true
			DC.Stubs = append(DC.Stubs[:i], DC.Stubs[i+1:]...)
			break
		}
	}
	fmt.Printf("Donator %d redeemed key stub\n", k.Data.Id)
	if !inArray {
		return ""
	}
	key, err := GenerateKey(ip, version, keyGenKey, true, DC.LastID, time.Now().Unix())
	DC.ActiveIDs = append(DC.ActiveIDs, DC.LastID)
	DC.LastID++
	writeConfig()
	return key
}

type (
	keyStub struct {
		Data struct {
			Id int `json:"id"`
		} `json:"data"`
		Time int64 `json:"time"`
	}
	DonatorConfig struct {
		ActiveIDs   []int     `json:"activeIDs"`
		InactiveIDs []int     `json:"inactiveIDs"`
		Stubs       []keyStub `json:"stubs"`
		LastID      int       `json:"lastID"`
		KeyStubID   int       `json:"keyStubID"`
	}
)
