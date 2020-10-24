package gash

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

func NewRSAToDir(dir string) error {
	private, err := os.OpenFile(dir+"/privateKey.pem", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer private.Close()
	public, err := os.OpenFile(dir+"/publicKey.pem", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer public.Close()
	return newKeyPair(private, public)
}

func NewRSAPair() (private, public []byte, err error) {
	var privateBuf bytes.Buffer
	var publicBuf bytes.Buffer
	if err = newKeyPair(&privateBuf, &publicBuf); err != nil {
		return
	}
	return privateBuf.Bytes(), publicBuf.Bytes(), err
}

func newKeyPair(pri, pub io.Writer) error {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	pemPrivateKey := pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509PrivateKey,
	}
	if err := pem.Encode(pri, &pemPrivateKey); err != nil {
		return err
	}

	publicKey := privateKey.PublicKey
	x509PublicKey, _ := x509.MarshalPKIXPublicKey(&publicKey)
	pemPublicKey := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509PublicKey,
	}
	if err := pem.Encode(pub, &pemPublicKey); err != nil {
		return err
	}
	return nil
}

type RSA struct {
	dir        string
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewRSAFromDir(dir string) (r *RSA, err error) {
	var privateKey, publicKey []byte
	if privateKey, err = ioutil.ReadFile(dir + "/private.pem"); err != nil {
		return nil, err
	}
	if publicKey, err = ioutil.ReadFile(dir + "/public.pem"); err != nil {
		return nil, err
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("publicKey key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	block, _ = pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("privateKey key error! ")
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return &RSA{
		dir:        dir,
		privateKey: pri,
		publicKey:  pub,
	}, nil
}

func (r *RSA) Encrypt(origin []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, origin)
}

func (r *RSA) Decrypt(cipherText []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, cipherText)
}

func (r *RSA) Sign(data []byte, h crypto.Hash) ([]byte, error) {
	hash := h.New()
	hash.Write(data)
	sign, err := rsa.SignPKCS1v15(rand.Reader, r.privateKey, h, hash.Sum(nil))
	if err != nil {
		return nil, err
	}
	return sign, nil
}

func (r *RSA) Verify(data []byte, sign []byte, h crypto.Hash) bool {
	hash := h.New()
	hash.Write(data)
	return rsa.VerifyPKCS1v15(r.publicKey, h, hash.Sum(nil), sign) == nil
}
