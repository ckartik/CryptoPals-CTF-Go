package util

import (
	"encoding/base64"
	"encoding/hex"
)

// FixedXOR xors the bytes of hstr1 with hstr2, for len(byte(hstr1)).
//
// Returns the resulting xor of hstr1 and hstr2 as a string.
func FixedXOR(hstr1, hstr2 string) string {
	b1, err := hex.DecodeString(hstr1)
	if err != nil {
		panic("There was an issue decoding the string")
	}
	b2, err := hex.DecodeString(hstr2)
	if err != nil {
		panic("There was an issue decoding the string")
	}

	XORValue := make([]byte, len(b1))
	for i := range b1 {
		XORValue[i] = b1[i] ^ b2[i]
	}

	return hex.EncodeToString(XORValue)
}

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

// Repeatedly xoring the bytestream of key with plaintext.
//
// Returns the ciphertext as a string
func DecryptRepeatingKeyXOR(plaintext, key string) string {

	// Convert Your input text into a sequence of bytes
	byteStream := []byte(plaintext)

	// Creates space for cipher.
	cipherStream := make([]byte, len(byteStream))

	// Initialize the key as bytes and coressponding params.
	byteKey := []byte(key)
	keySize := len(byteKey)

	// This is like auto in C++.
	for i := range byteStream {
		cipherStream[i] = byteStream[i] ^ byteKey[i%keySize]
	}

	return string(cipherStream)
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
