package main

import (
	"./common/hammingdist"
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"os"
)

// Frequency attack against a single charecter XOR.
func singleByteXOR(hexcipher string) string {
	byteStream, err := hex.DecodeString(hexcipher)
	if err != nil {
		panic("There was an issue decoding the string")
	}

	// Our current best guess as to the plaintext.
	plaintext := string(byteStream)
	bestScore := hammingdist.PlaintextScore(plaintext)

	// We will take the L2 Norm between charecter frequence of english language and set given.
	XORValue := make([]byte, len(byteStream))
	for j := 0x0; j < 0xff; j++ { // Keep top scorer only, otherwise too memory expensive.
		for i := range byteStream {
			XORValue[i] = byteStream[i] ^ byte(j+0x1)
		}
		guess := string(XORValue)
		score := hammingdist.PlaintextScore(guess)
		if score < bestScore {
			plaintext = guess
			bestScore = score
		}
	}

	return plaintext
}

func detectSingleKeyXOR(file *File) string {
	file, err := os.Open("4.txt")
	check(err)

	defer file.Close()

	bestScore := 100.0
	bestString := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := singleByteXOR(scanner.Text())
		score := hammingdist.PlaintextScore(text)
		if score < bestScore {
			bestScore = score
			bestString = text
		}
	}
	check(scanner.Err())

	return bestString
}

func main() {

	// Local Test of S1C1
	S1C3Input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	S1C3Answer := "Cooking MC's like a pound of bacon"
	S1C3Result := singleByteXOR(S1C3Input)
	if S1C3Result != S1C3Answer {
		fmt.Printf(S1C3Result)
	}

	S1C4Result := detectSingleKeyXOR()
	S1C4Answer := "Now that the party is jumping\n"
	if S1C4Result != S1C4Answer {
		fmt.Printf(S1C4Result)
	}

}
