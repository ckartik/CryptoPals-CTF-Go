package util

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"log"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Takes in hex-encoded ciphers.
func DetectAES(ciphers <-chan []byte) chan string {
	done := make(chan string, 1)

	go func(ciphers <-chan []byte, done chan string) {
		for cipher := range ciphers {
			decoded := make([]byte, len(cipher)/2)
			hmap := make(map[string]bool)
			n, err := hex.Decode(decoded, cipher)
			check(err)

			for index := 0; index <= n-16; index += 16 {
				word := string(cipher[index : index+16])

				if hmap[word] {
					done <- string(cipher)

					return
				}

				hmap[word] = true
			}

		}

		done <- ""
	}(ciphers, done)

	return done
}

// DecryptAES takes in a byte-array of cipher and a key and decryptes it.
//
// Does not currently through an error but should.
//
// TODO: Design question - change cipher byte-array to io.Reader and return an io.reader as well.
//	 since I handle breaking up the cipher bytes array into seperate blocks and I don't need to
//	 parralize the decrypt operations.
func DecryptAES(cipher []byte, key []byte) string {
	cblock, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
	}

	cipherSize := len(cipher)
	keySize := len(key)

	plaintext := make([]byte, cipherSize)

	// TODO: Read only multiple of keysize and throw error if filesize is corrupted and not a multiple.
	var j int
	for j = 0; j < cipherSize; j += keySize {
		cblock.Decrypt(plaintext[j:], cipher[j:])
	}

	return string(plaintext)
}

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

const BlockSize8 int = 8
const BlockSize16 int = 16

// TODO: Support only certain encodings.
// TODO: Add encoding detection and fail if not correctly encoded.
func PKCS7(plaintext []byte, blockSize int) []byte {
	originalSize := len(plaintext)
	paddingSize := blockSize - (originalSize % blockSize)

	padding := make([]byte, paddingSize)

	for i := range padding {
		padding[i] = byte(4)
	}

	return append(plaintext[:], padding[:]...)
}
