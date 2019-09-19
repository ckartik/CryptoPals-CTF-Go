package util

import (
	"encoding/base64"
	"encoding/hex"
)

// Repeatedly xoring the bytestream of key with plaintext.
//
// Returns the ciphertext as a string
func RepeatingKeyXOR(plaintext, key string) string {
	byteStream := []byte(plaintext)
	cipherStream := make([]byte, len(byteStream))
	byteKey := []byte(key)
	keySize := len(byteKey)
	for i := range byteStream {
		cipherStream[i] = byteStream[i] ^ byteKey[i%keySize]
	}

	return hex.EncodeToString(cipherStream)
}

// Hexto64 re-encodes the passed in hexString as base64.
//
// Returns it as a base64 encoding.
func HexTo64(hexString string) string {
	binaryData, err := hex.DecodeString(hexString)
	if err != nil {
		panic("There was an issue decoding the string")
	}

	return base64.StdEncoding.EncodeToString(binaryData)
}
