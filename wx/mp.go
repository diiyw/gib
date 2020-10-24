package wx

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
)

type MP struct {
	OpenID    string `json:"openid"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
	Country   string `json:"country"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Language  string `json:"language"`
	Province  string `json:"province"`
	Watermark struct {
		AppID string `json:"appid"`
	} `json:"watermark"`
}

// Decrypt WeiXin APP's AES Data
func BizDataDecrypt(encryptedData, iv, appId, sessionKey string) (*MP, error) {
	if len(sessionKey) != 24 {
		return nil, errors.New("session key length is error ")
	}
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	if len(iv) != 24 {
		return nil, errors.New("iv length is error")
	}
	aesIv, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	aesCipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	aesPlaintext := make([]byte, len(aesCipherText))

	aesBlock, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(aesBlock, aesIv)
	mode.CryptBlocks(aesPlaintext, aesCipherText)
	aesPlaintext = PKCS7UnPadding(aesPlaintext)

	var mp *MP
	err = json.Unmarshal(aesPlaintext, &mp)
	if err != nil {
		return nil, err
	}

	if mp.Watermark.AppID != appId {
		return nil, errors.New("appId is not match")
	}

	return mp, nil
}

// PKCS7UnPadding return unPadding []Byte plaintext
func PKCS7UnPadding(plaintext []byte) []byte {
	length := len(plaintext)
	if length > 0 {
		unPadding := int(plaintext[length-1])
		return plaintext[:(length - unPadding)]
	}
	return plaintext
}
