package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// This should be in an env file in production
const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func Encode(b []byte) (string, error) {
	return base64.StdEncoding.EncodeToString(b), nil
}
func Decode(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return []byte{}, nil
	}
	return data, nil
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	enc, err := Encode(cipherText)
	if err != nil {
		return "", err
	}
	return enc, nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText, err := Decode(text)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

type CreateSessionParams struct {
	LdapURL      string `json:"ldap_url" bson:"ldap_url" binding:"required"`
	BindDN       string `json:"bind_dn" bson:"bind_dn" binding:"required"`
	BindPassword string `json:"bind_password" bson:"bind_password" binding:"required"`
}

func CreateSession(Params *CreateSessionParams) (string, error) {
	text := Params.LdapURL + "|" + Params.BindDN + "|" + Params.BindPassword
	enctext, err := Encrypt(text, MySecret)
	if err != nil {
		return "", err
	}
	return enctext, nil
}

func ReadSession(sessionHash string) (*CreateSessionParams, error) {
	dec, err := Decrypt(sessionHash, MySecret)
	if err != nil {
		return nil, err
	}
	decArr := strings.Split(dec, "|")
	if len(decArr) != 3 {
		return nil, fmt.Errorf("Invalid session hash")
	}
	if err != nil {
		return nil, err
	}
	return &CreateSessionParams{
		LdapURL:      decArr[0],
		BindDN:       decArr[1],
		BindPassword: decArr[2],
	}, nil
}
