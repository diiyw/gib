package gash

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(in string) string {
	h := md5.New()
	h.Write([]byte(in))
	return hex.EncodeToString(h.Sum(nil))
}
