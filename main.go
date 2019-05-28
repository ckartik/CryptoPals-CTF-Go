package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
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

var englishModel = []float64{
	0.0651738, 0.0124248, 0.0217339, 0.0349835, //'A', 'B', 'C', 'D',...
	0.1041442, 0.0197881, 0.0158610, 0.0492888,
	0.0558094, 0.0009033, 0.0050529, 0.0331490,
	0.0202124, 0.0564513, 0.0596302, 0.0137645,
	0.0008606, 0.0497563, 0.0515760, 0.0729357,
	0.0225134, 0.0082903, 0.0171272, 0.0013692,
	0.0145984, 0.0007836, 0.1918182}

const (
	captialStart = 0x61
	captialEnd   = 0x7a
	lowerStart   = 0x41
	lowerEnd     = 0x5a
)

// manhatanDistance finds the manhatan distance between vector1 and vector2.
func pnorm(vector1, vector2 []float64, p float64) float64 {
	dist := 0.0
	for i := range vector1 {
		dist += math.Pow(math.Abs(vector1[i]-vector2[i]), p)
	}
	return math.Pow(dist, 1/p)
}

// plaintextscore finds the manhatan distance between text and the english model.
func plaintextScore(text string) float64 {
	stringVector := make([]float64, len(englishModel))
	for i := range text {
		hex := byte(text[i])
		if hex >= captialStart && hex <= captialEnd {
			stringVector[int(hex-captialStart)]++
		} else if hex >= lowerStart && hex <= lowerEnd {
			stringVector[int(hex-lowerStart)]++
		}
	}

	plaintextSize := len(text)
	for i := range stringVector {
		stringVector[i] /= float64(plaintextSize)
	}

	return pnorm(stringVector, englishModel, .5)
}

// Frequency attack against a single charecter XOR.
func singleByteXOR(hexcipher string) string {
	byteStream, err := hex.DecodeString(hexcipher)
	if err != nil {
		panic("There was an issue decoding the string")
	}

	// Our current best guess as to the plaintext.
	plaintext := string(byteStream)
	bestScore := plaintextScore(plaintext)

	// We will take the L2 Norm between charecter frequence of english language and set given.
	XORValue := make([]byte, len(byteStream))
	for j := 0x0; j < 0xff; j++ { // Keep top scorer only, otherwise too memory expensive.
		for i := range byteStream {
			XORValue[i] = byteStream[i] ^ byte(j+0x1)
		}
		guess := string(XORValue)
		score := plaintextScore(guess)
		if score < bestScore {
			plaintext = guess
			bestScore = score
		}
	}

	return plaintext
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
		fmt.Printf(S1C3Result)
	}
}
