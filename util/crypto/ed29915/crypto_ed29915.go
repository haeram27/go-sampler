package ed29915

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

// Asymmetric Key - ED25519

type ED25519Helper struct {
	PrivateKey *ed25519.PrivateKey
	PublicKey  *ed25519.PublicKey
}

/**
*	Generate
**/
func (e *ED25519Helper) Generate() error {

	if e.PrivateKey != nil {
		e.PrivateKey = nil
		e.PublicKey = nil
	}

	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)

	if err != nil {
		return err
	}

	e.PrivateKey = &privateKey
	e.PublicKey = &publicKey
	return nil
}

/**
*	EncodePrivateKey
**/
func (e *ED25519Helper) EncodePrivateKey() ([]byte, error) {
	if e.PrivateKey == nil {
		return nil, errors.New("private key is nil")
	}

	encBuf := pem.EncodeToMemory(
		&pem.Block{
			Type:  "ED25519 PRIVATE KEY",
			Bytes: *e.PrivateKey,
		})

	return encBuf, nil
}

/**
*	DecodePrivateKey
**/
func (e *ED25519Helper) DecodePrivateKey(encKey []byte) error {
	if e.PrivateKey != nil {
		e.PrivateKey = nil
	}

	block, _ := pem.Decode(encKey)

	if block == nil {
		return errors.New("failed to find ED25519 PRIVATE KEY")
	}

	if block.Type != "ED25519 PRIVATE KEY" {
		return fmt.Errorf("%s is invalid block type", block.Type)
	}
	e.PrivateKey = (*ed25519.PrivateKey)(&block.Bytes)
	return nil
}

/**
*	PrivateKeyToFile
**/
func (e *ED25519Helper) PrivateKeyToFile(fname string) error {
	if e.PrivateKey == nil {
		return errors.New("private key is nil")
	}
	enc, err := e.EncodePrivateKey()

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fname, enc, 0644)

	if err != nil {
		return err
	}
	return nil
}

/**
*	PrivateKeyFromFile
**/
func (e *ED25519Helper) PrivateKeyFromFile(fname string) error {
	if e.PrivateKey != nil {
		e.PrivateKey = nil
	}

	buf, err := ioutil.ReadFile(fname)

	if err != nil {
		return err
	}

	err = e.DecodePrivateKey(buf)

	if err != nil {
		return err
	}
	return nil
}

/**
*	EncodePublicKey
**/
func (e *ED25519Helper) EncodePublicKey() ([]byte, error) {
	if e.PublicKey == nil {
		return nil, errors.New("public key is nil")
	}

	encBuf := pem.EncodeToMemory(
		&pem.Block{
			Type:  "ED25519 PUBLIC KEY",
			Bytes: *e.PublicKey,
		})

	return encBuf, nil
}

/**
*	DecodePublicKey
**/
func (e *ED25519Helper) DecodePublicKey(encKey []byte) error {
	if e.PublicKey != nil {
		e.PublicKey = nil
	}

	block, _ := pem.Decode(encKey)

	if block == nil {
		return errors.New("failed to find ED25519 PUBLIC KEY")
	}

	if block.Type != "ED25519 PUBLIC KEY" {
		return fmt.Errorf("%s is invalid block type", block.Type)
	}

	e.PublicKey = (*ed25519.PublicKey)(&block.Bytes)

	return nil
}

/**
*	PublicKeyToFile
**/
func (e *ED25519Helper) PublicKeyToFile(fname string) error {
	if e.PublicKey == nil {
		return errors.New("public key is nil")
	}
	enc, err := e.EncodePublicKey()

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fname, enc, 0644)

	if err != nil {
		return err
	}
	return nil
}

/**
*	PublicKeyFromFile
**/
func (e *ED25519Helper) PublicKeyFromFile(fname string) error {
	if e.PublicKey != nil {
		e.PrivateKey = nil
	}

	buf, err := ioutil.ReadFile(fname)

	if err != nil {
		return err
	}

	err = e.DecodePublicKey(buf)

	if err != nil {
		return err
	}
	return nil
}

/**
*	Sign
**/
func (e *ED25519Helper) Sign(msg string) ([]byte, error) {
	if e.PrivateKey == nil {
		return nil, errors.New("private key is nil")
	}

	hash := sha256.Sum256([]byte(msg))
	signed := ed25519.Sign(*e.PrivateKey, hash[:])

	return signed, nil
}

/**
*	SignToString
**/
func (e *ED25519Helper) SignToString(msg string) (string, error) {
	signed, err := e.Sign(msg)

	if err != nil {
		return "", err
	}
	return hex.EncodeToString(signed), nil
}

/**
* SignToPEM
**/
func (e *ED25519Helper) SignToPEM(msg string, blockType string) ([]byte, error) {
	signed, err := e.Sign(msg)

	if err != nil {
		return nil, err
	}

	encBuf := pem.EncodeToMemory(
		&pem.Block{
			Type:  blockType,
			Bytes: signed,
		})

	return encBuf, nil
}

/**
*	SignFromPEM
**/
func (e *ED25519Helper) SignFromPEM(pemBuf []byte, blockType string) ([]byte, error) {
	block, _ := pem.Decode(pemBuf)

	if block == nil {
		return nil, errors.New("failed to find ED25519 PUBLIC KEY")
	}

	if block.Type != blockType {
		return nil, fmt.Errorf("%s is invalid block type", block.Type)
	}

	return block.Bytes, nil
}

/**
*	IsValidSignPEM
**/
func (e *ED25519Helper) IsValidSignPEM(pemBuf []byte, blockType string) bool {
	block, _ := pem.Decode(pemBuf)

	if block == nil {
		return false
	}

	if block.Type != blockType {
		return false
	}
	return true
}

/**
*	Verify
**/
func (e *ED25519Helper) Verify(msg string, signed []byte) bool {
	if e.PublicKey == nil {
		return false
	}
	hash := sha256.Sum256([]byte(msg))

	return ed25519.Verify(*e.PublicKey, hash[:], signed)
}

/**
*	VerifyString
**/
func (e *ED25519Helper) VerifyString(msg string, signed string) bool {
	byteSigned, err := hex.DecodeString(signed)

	if err != nil {
		return false
	}

	return e.Verify(msg, byteSigned)
}
