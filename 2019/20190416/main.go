package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	fmt.Println(AESCBCEncrypter("exampleplaintext你妈炸sdfjhvsdfhbn((&&^$$%^#$%^jvsjbdfvbhj了"))
	fmt.Println(AESCBCDecrypter("f24c318676c2c85aca442453c5c5e0c53dcd175db4b71c17d6a5bdc58c2371a500cb3b8bfbd5c6495f442a986303574af29deb83cb6f6f93aa974e018940d4349bd900a9b9cd4de3d3499e0e923dc76849a078a55a290193cd143bb3f1e6de4f981c3c90a3b76689f4054cbead7a147fab035b99b61832df2d2c9afcebfb9991e20017abc192f4bf20b9e534e6d7ca1bf28dfcff119fec6d13b8390fd10e85f5"))
}

func AESCBCDecrypter(cipher_text string) string {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	ciphertext, _ := hex.DecodeString(cipher_text)
	block, err := aes.NewCipher(key)
	if err != nil {
		return cipher_text
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = PKCS5UnPadding(ciphertext)
	return fmt.Sprintf("%s", ciphertext)
}

func AESCBCEncrypter(plain_text string) string {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte(plain_text)
	plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	block, err := aes.NewCipher(key)
	if err != nil {
		return plain_text
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return plain_text
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return fmt.Sprintf("%x", ciphertext)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
