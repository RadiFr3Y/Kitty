package files

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

// Encrypt code
func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// calculate Encrypt
func Encrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	Crypt := []byte(text)
	Cfb := cipher.NewCFBEncrypter(block, bytes)
	Text := make([]byte, len(Crypt))
	Cfb.XORKeyStream(Text, Crypt)
	return Encode(Text), nil
}
