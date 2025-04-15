package jws

import (
	"bytes"
	creek_sdk "creek-sdk"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"gopkg.in/go-jose/go-jose.v2"
)

func NewRandomSymmetricKey(size int) ([]byte, error) {
	if size <= 0 || size%8 != 0 {
		return nil, creek_sdk.WithStack(fmt.Errorf("invalid key size"))
	}

	k := make([]byte, size)
	if _, err := rand.Read(k); err != nil {
		return nil, creek_sdk.WithStack(err)
	}
	return k, nil
}

func NewRandomRSAKey() (*rsa.PrivateKey, error) {
	// Generate a public/private key pair to use for this example.
	return rsa.GenerateKey(rand.Reader, 4096)
}

func ExportPrivateKey(pk *rsa.PrivateKey) ([]byte, error) {
	var pemPrivateBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pk),
	}
	buffer := new(bytes.Buffer)
	if err := pem.Encode(buffer, pemPrivateBlock); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func ExportPublicKey(pk *rsa.PrivateKey) ([]byte, error) {
	var pemPublicBlock = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&pk.PublicKey),
	}
	buffer := new(bytes.Buffer)
	if err := pem.Encode(buffer, pemPublicBlock); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func NewPublicKeyFromPEM(pk []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pk)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	}
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pubKey, nil
}

func NewPrivateKeyFromPEM(pk []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pk)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to parse PEM block containing the private key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func NewSigner(privateKey *rsa.PrivateKey) (jose.Signer, error) {
	return jose.NewSigner(jose.SigningKey{Algorithm: jose.PS512, Key: privateKey}, nil)
}

func NewHMACSigner(secret []byte) (jose.Signer, error) {
	sign, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS512, Key: secret}, nil)
	if err != nil {
		return nil, creek_sdk.WithStack(err)
	}
	return sign, nil
}
