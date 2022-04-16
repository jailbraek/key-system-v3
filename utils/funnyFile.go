package utils

import (
	"encoding/base64"
	"github.com/thanhpk/randstr"
)

func FunnyEncoding(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

func FunnyDecoding(data string) ([]byte, error) {
	emptyArry := make([]byte, 0)
	cipherText, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return emptyArry, err
	}
	return cipherText, nil
}

func GenerateBrowserID() string {
	return FunnyEncoding(randstr.Bytes(100))
}
