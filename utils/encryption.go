package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func Encrypt(data []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return "", err
	}
	ciphertext := aesGCM.Seal(nonce, nonce, data, []byte("PENISPENISPENISDARKHUBBEST"))
	return FunnyEncoding(ciphertext), nil
}

func Decrypt(text string, key []byte) (string, error) {
	cipherText, err := FunnyDecoding(text)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	aesGCM, err := cipher.NewGCM(block)
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, []byte("PENISPENISPENISDARKHUBBEST"))
	return string(plaintext), err
}
