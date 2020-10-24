package gash

import (
	"testing"
)

func TestNew(t *testing.T) {
	if err := NewRSAToDir("testdata"); err != nil {
		t.Error(err)
	}
}

func TestRSA(t *testing.T) {
	r, err := NewRSAFromDir("testdata")
	if err != nil {
		t.Error(err)
	}
	encrypted, err := r.Encrypt([]byte("hello rsa"))
	if encrypted == nil {
		t.Failed()
	}
	b, err := r.Decrypt(encrypted)
	if err != nil {
		t.Error(err)
	}
	if string(b) != "hello rsa" {
		t.Failed()
	}
}
