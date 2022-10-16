package files

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// clear console function
func Clear() {
	system := runtime.GOOS
	if system == "windows" {
		console := exec.Command("cmd", "/c", "cls")
		console.Stdout = os.Stdout
		console.Run()
	} else {
		console := exec.Command("clear")
		console.Stdout = os.Stdout
		console.Run()
	}
}

// Decrypt code
func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		Clear()
		fmt.Println("Error!")
	}
	return data
}

// Calculate Decrypt
func Decrypt(text, MySecret string) (string, error) {
	data, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		Clear()
		fmt.Println("Error!")
	}
	Crypt := Decode(text)
	Cfb := cipher.NewCFBDecrypter(data, bytes) // "bytes" variable is in Encrypt.go File
	Text := make([]byte, len(Crypt))
	Cfb.XORKeyStream(Text, Crypt)
	return string(Text), nil
}
