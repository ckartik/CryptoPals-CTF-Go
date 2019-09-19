package attack

import (
	"../../common/hammingdist"
	"bufio"
	"encoding/hex"
	"os"
)

// Frequency attack against a single charecter XOR over a encrrpted text file.
//
// returns the best perdicted plaintext.
func DetectSingleKeyXOR(file *os.File) string {
	defer file.Close()

	bestScore := 100.0
	bestString := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text, _ := singleByteXOR(scanner.Text())
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

// Frequency attack against a single charecter XOR. Attack works by brute-forcing every key against
// a similarty measure of charecter frequencies using the hammingdist plaintext score.
// returns the tuple plaintext, key.
func singleByteXOR(hexcipher string) (string, byte) {
	byteStream, err := hex.DecodeString(hexcipher)
	if err != nil {
		panic("There was an issue decoding the string")
	}

	// Our current best guess as to the plaintext.
	plaintext := string(byteStream)
	bestScore := hammingdist.PlaintextScore(plaintext)

	// We will take the L2 Norm between charecter frequence of english language and set given.
	XORValue := make([]byte, len(byteStream))
	var key byte
	for j := 0x0; j < 0xff; j++ { // Keep top scorer only, otherwise too memory expensive.
		for i := range byteStream {
			XORValue[i] = byteStream[i] ^ byte(j+0x1)
		}
		guess := string(XORValue)
		score := hammingdist.PlaintextScore(guess)
		if score < bestScore {
			plaintext = guess
			bestScore = score
			key = byte(j + 0x1)
		}
	}

	return plaintext, key
}
