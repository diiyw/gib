package hash

import (
	"fmt"
	"testing"
	"time"
)

func TestDesEncrypt(t *testing.T) {
	d, err := DesEncrypt([]byte("hello"), []byte("12345678"), time.Now().Add(time.Hour*12))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("DesEncrypt:", string(d))
	s, err := DesDecrypt(d, []byte("12345678"))
	if string(s) != "hello" {
		t.Fatal("decrypt error:", err)
	}
}
