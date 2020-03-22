package hash

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/diiyw/gib/text"
	"time"
)

var (
	ContentError = errors.New("content error")
	Timeout      = errors.New("timeout")
)

func MD5(in string) string {
	h := md5.New()
	h.Write([]byte(in))
	return hex.EncodeToString(h.Sum(nil))
}

func DesEncrypt(originData []byte, key []byte, expired time.Time) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	originData = append(originData, '-')
	originData = append(originData, []byte(expired.Format(text.DateFormat))...)
	originData = PKCS5Padding(originData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	cryptData := make([]byte, len(originData))
	blockMode.CryptBlocks(cryptData, originData)
	return []byte(base64.StdEncoding.EncodeToString(cryptData)), nil

}

func DesDecrypt(cryptData []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	cryptData, err = base64.StdEncoding.DecodeString(string(cryptData))
	if err != nil {
		return nil, err
	}
	originData := make([]byte, len(cryptData))
	blockMode.CryptBlocks(originData, cryptData)
	originData = PKCS5UnPadding(originData)
	bb := bytes.SplitN(originData, []byte("-"), 2)
	if len(bb) < 2 {
		return nil, ContentError
	}
	t, err := time.Parse(text.DateFormat, string(bb[1]))

	if err != nil {
		return nil, err
	}
	if t.Before(time.Now()) {
		return nil, Timeout
	}
	return bb[0], nil

}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
