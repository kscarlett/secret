package encryption

import (
	"testing"
)

const (
	key           = "testPassword01"
	data          = "TestSuccess"
	encryptedData = "cb93cdf3f8aa99a2a4a05b8fdc3cacb4ece1abf0d2260271ae9309"
)

func TestEncryption(t *testing.T) {
	_, err := Encrypt(key, data)
	if err != nil {
		t.Fatalf("failed to encrypt data - error returned: %s", err.Error())
	}

	// TODO: check if there is a way to verify that the encrypted data makes sense
}

func TestDecryption(t *testing.T) {
	decrypted, err := Decrypt(key, encryptedData)
	if err != nil {
		t.Fatalf("failed to decrypt data - error returned: %s", err.Error())
	}

	if decrypted != data {
		t.Errorf("failed to decrypt data - got '%s', expected '%s'", decrypted, data)
	}
}

func TestEncryptThenDecrypt(t *testing.T) {
	encryptedString, err := Encrypt(key, data)
	if err != nil {
		t.Fatalf("failed to encrypt data - error returned: %s", err.Error())
	}

	decryptedString, err := Decrypt(key, encryptedString)
	if err != nil {
		t.Fatalf("failed to decrypt data - error returned: %s", err.Error())
	}

	if decryptedString != data {
		t.Errorf("decrypted data does not match encrypted data - got '%s', expected '%s'", decryptedString, data)
	}
}

func BenchmarkEncryptThenDecrypt(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		encryptedString, err := Encrypt(key, string(i))
		if err != nil {
			b.Fatalf("failed to encrypt data - error returned: %s", err.Error())
		}

		decryptedString, err := Decrypt(key, encryptedString)
		if err != nil {
			b.Fatalf("failed to decrypt data - error returned: %s", err.Error())
		}

		if decryptedString != string(i) {
			b.Errorf("decrypted data does not match encrypted data - got '%s', expected '%s'", decryptedString, string(i))
		}
	}

}
