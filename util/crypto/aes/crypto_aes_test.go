package aes

import (
	"testing"
)

// Symmetric Key - AES

func TestAES(t *testing.T) {

	// cipher key,  32 bytes to select AES-256.
	key := "thisis32bytlongpassphraseimusing"

	// plaintext
	pt := "This is a secret"

	c := EncryptAES([]byte(key), pt)

	// plaintext
	t.Log(pt)

	// ciphertext
	t.Log(c)

	// decrypt
	d := DecryptAES([]byte(key), c)

	// decrypted text
	t.Log(d)
}
