package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

type Key struct {
	Ip        string `json:"ip"`
	Time      int64  `json:"time"`
	Donator   bool   `json:"donator"`
	DonatorID int    `json:"donatorID,omitempty"`
	Version   string `json:"version"`
}

func GenerateKey(ip, version string, keyGenKey []byte, donator bool, donatorID int, time int64) (string, error) {
	key := Key{
		Ip:      ip,
		Time:    time,
		Donator: donator,
		Version: version,
	}
	if donator {
		key.DonatorID = donatorID
	}
	keyData, err := json.Marshal(key)
	if err != nil {
		return "", err
	}
	encrypt, err := Encrypt(keyData, keyGenKey)
	if err != nil {
		return "", err
	}
	return encrypt, nil
}

func ParseKey(key string, keyGenKey []byte) (Key, error) {
	keyData, err := Decrypt(key, keyGenKey)
	if err != nil {
		return Key{}, err
	}
	var keyStruct Key
	err = json.Unmarshal([]byte(keyData), &keyStruct)
	if err != nil {
		return Key{}, err
	}
	return keyStruct, nil
}

func CheckKey(key, ip, version string, keyGenKey []byte) (bool, error) {
	keyStruct, err := ParseKey(key, keyGenKey)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if keyStruct.Ip != ip {
		return false, nil
	}
	if keyStruct.Donator {
		loadConfig()
		if keyStruct.DonatorID >= 0 {
			for _, id := range DC.ActiveIDs {
				if keyStruct.DonatorID == id {
					return true, nil
				}
			}
		}
		return false, nil
	}
	if keyStruct.Version != version {
		return false, nil
	}
	if keyStruct.Time < time.Now().Unix() {
		return false, nil
	}
	return true, nil
}
