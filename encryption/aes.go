package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

//AesEncryptionService encrypts value using AES
//Don't create it directly! Use constructor NewAesEncryptionService instead!
type AesEncryptionService struct {
	secret []byte
}

//NewAesEncryptionService creates AesEncryptionService structure.
func NewAesEncryptionService(secret string) (*AesEncryptionService, error) {
	// todo: do smth with `32`
	if len(secret) != 32 {
		return nil, fmt.Errorf("secret should be 32 byte length")
	}

	return &AesEncryptionService{secret: []byte(secret)}, nil
}

//Encrypt encrypts some []byte data
func (service AesEncryptionService) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(service.secret)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText, nil
}

//Decrypt decrypts some []byte data
func (service AesEncryptionService) Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(service.secret)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
