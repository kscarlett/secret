package secret

import (
	"errors"

	"github.com/kscarlett/secret/pkg/encryption"
)

func NewInMemory(encryptionKey string) Secret {
	return Secret{
		encryptionKey: encryptionKey,
		keyValues:     make(map[string]string),
	}
}

type Secret struct {
	encryptionKey string
	keyValues     map[string]string
}

func (s *Secret) Get(key string) (string, error) {
	encryptedValue, ok := s.keyValues[key]
	if !ok {
		return "", errors.New("no value found for the given key")
	}

	ret, err := encryption.Decrypt(s.encryptionKey, encryptedValue)
	if err != nil {
		return "", err
	}

	return ret, nil
}

func (s *Secret) Set(key, value string) error {
	encryptedValue, err := encryption.Encrypt(s.encryptionKey, value)
	if err != nil {
		return err
	}

	s.keyValues[key] = encryptedValue
	return nil
}
