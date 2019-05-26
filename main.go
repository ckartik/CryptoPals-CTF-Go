package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// hexto64 re-encodes the passed in hexString and returns it as a base64 encoding.
func hexto64(hexString string) string {
	binaryData, err := hex.DecodeString(hexString)
	if err != nil {
		panic("There was an issue decoding the string")
	}

	return base64.StdEncoding.EncodeToString(binaryData)
}

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

// Frequency attack against a single charecter XOR.
func singleByteXOR(hexcipher string) string {
	byteStream, err := hex.DecodeString(hexcipher)
	if err != nil {
		panic("There was an issue decoding the string")
	}

	// Our current best guess as to the plaintext.
	plaintext := string(byteStream)

	// We will take the L2 Norm between charecter frequence of english language and set given.
	XORValue := make([]byte, len(byteStream))
	for j := 0x0; j < 0xff; j++ { // Keep top scorer only, otherwise too memory expensive.
		for i := range byteStream {
			XORValue[i] = byteStream[i] ^ byte(j+0x1)
		}
		potentialValue := string(XORValue)
	}

	return string(byteStream)
}

func main() {
	// Local Test of S1C1
	S1C1Input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	S1C1Answer := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	S1C1Result := hexto64(S1C1Input)
	if S1C1Result != S1C1Answer {
		fmt.Printf(S1C1Result)
	}

	// Local Test of S1C2
	S1C2Input := "1c0111001f010100061a024b53535009181c"
	S1C2Input2 := "686974207468652062756c6c277320657965"
	S1C2Answer := "746865206b696420646f6e277420706c6179"
	S1C2Result := FixedXOR(S1C2Input, S1C2Input2)
	if S1C2Result != S1C2Answer {
		fmt.Printf(S1C2Result)
	}

	// Local Test of S1C1
	S1C3Input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	S1C3Answer := "Cooking MC's like a pound of bacon"
	S1C3Result := singleByteXOR(S1C3Input)
	if S1C3Result != S1C3Answer {
		fmt.Printf(S1C1Result)
	}
}
