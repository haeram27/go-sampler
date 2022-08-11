package aes

import (
	"crypto/aes"
	"encoding/hex"
)

// Symmetric Key - AES
/*
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
*/

func EncryptAES(key []byte, plaintext string) string {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	return s
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
