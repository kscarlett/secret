package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// Encrypt encrypts the given data using AES-256
// The data is returned as a hex string.
// This code is based on the example proviced in the official documentation.
func Encrypt(key, data string) (string, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(data))

	return fmt.Sprintf("%x", ciphertext), nil
}

// Decrypt decrypts the given data using AES-256. The data is passed in as a hex string.
// This code is based on the example proviced in the official documentation.
func Decrypt(key, data string) (string, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("encryption: ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := sha256.New()
	io.WriteString(hasher, key)
	keyHash := hasher.Sum(nil)
	return aes.NewCipher(keyHash)
}
