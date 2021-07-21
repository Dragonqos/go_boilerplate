package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

const BlockSize32 = aes.BlockSize*2

func Encrypt(key []byte, iv []byte, text string) string {
	pkey := pkcs5Pad(key, BlockSize32)
	block, err := aes.NewCipher(pkey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	byteText := []byte(text)
	byteText = pkcs5Pad(byteText, aes.BlockSize)

	encrypted := make([]byte, len(byteText))
	mode.CryptBlocks(encrypted, byteText)

	return base64.RawStdEncoding.EncodeToString(encrypted)
}

func Decrypt(key []byte, iv []byte, cryptoText string) string {
	pkey := pkcs5Pad(key, BlockSize32)
	block, err := aes.NewCipher(pkey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	byteText, err1 := base64.RawStdEncoding.DecodeString(cryptoText)
	if err1 != nil {
		panic(err1)
	}

	if len(byteText)%mode.BlockSize() != 0 {
		byteText = pkcs5Pad(byteText, aes.BlockSize)
	}

	decrypted := make([]byte, len(byteText))
	mode.CryptBlocks(decrypted, byteText)

	return string(pkcs5Trim(decrypted))
}

func pkcs5Pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5Trim(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}