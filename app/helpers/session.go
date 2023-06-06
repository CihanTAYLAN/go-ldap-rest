package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// This should be in an env file in production
const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		log.Fatal(err)
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		log.Fatal(err)
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

type CreateSessionParams struct {
	LdapURL      string `json:"ldapURL" bson:"searchBase" binding:"required"`
	BindDN       string `json:"bindDN" bson:"searchFilter" binding:"required"`
	BindPassword string `json:"bindPassword" bson:"attributes" binding:"required"`
}

func CreateSession(Params *CreateSessionParams) (string, error) {
	text := Params.LdapURL + "|" + Params.BindDN + "|" + Params.BindPassword
	enctext, err := Encrypt(text, MySecret)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return enctext, nil
}

func ReadSession(sessionHash string) (*CreateSessionParams, error) {
	dec, err := Decrypt(sessionHash, MySecret)
	print(dec)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return nil, nil
}
