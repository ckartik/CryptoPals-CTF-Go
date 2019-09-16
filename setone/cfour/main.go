package main

import (
	"../../common/hammingdist"
	"bufio"
	"encoding/hex"
	"fmt"
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

func detectSingleKeyXOR(file *os.File) string {
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

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return bestString
}

func main() {
	file, err := os.Open("4.txt")
	if err != nil {
		panic(err)
	}
	S1C4Result := detectSingleKeyXOR(file)
	S1C4Answer := "Now that the party is jumping\n"
	if S1C4Result != S1C4Answer {
		fmt.Printf(S1C4Result)
	}
}
