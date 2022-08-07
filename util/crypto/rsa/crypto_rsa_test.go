package rsa

import (
	"testing"
)

// Asymmetric Key - RSA

func TestRsa(t *testing.T) {
	rh := RSAHelper{}
	rh.GenerateKey(1024)

	priv_pem, _ := rh.PrivateToStringPEM()
	pub_pem, _ := rh.PublicToStringPEM()

	t.Log(priv_pem)
	t.Log(pub_pem)

	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*())-_=+"

	t.Log("plain", s)

	enc, _ := rh.EncryptString(s)

	t.Log(enc)

	dec, _ := rh.DecryptString(enc)

	t.Log(dec)
}
