package attack

import (
	"../common/hammingdist"
	"../common/util"
	"bufio"
	"encoding/hex"
	"log"
	"math"
	"os"
)

// TODO: Break up into Returning Keysize function and creating blocks, only pass a slice of the cipher into this and all other funcitions.
func breakRepeatingKeyXOR() string {
	file, err := os.Open("6.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	cipher := make([]byte, 4000)
	bytesRead, err := file.Read(cipher)
	if err != nil {
		panic(err)
	} else {
		log.Printf("Successfully Read in %v bytes", bytesRead)
	}

	guesses := hammingdist.GuessKeySize(cipher[:])
	bestScore := math.Inf(1)
	var bestPlaintext string
	for i := 0; i < 5; i++ {
		bestKeySize := []hammingdist.Guess(guesses)[i].Keysize

		log.Printf("Found a perdicted keysize of %v", bestKeySize)

		// Break the ciphertext into blocks of keySize length.
		numOfBlocks := int(math.Ceil(float64(bytesRead) / float64(bestKeySize)))
		blocks := make([][]byte, numOfBlocks)
		base := 0
		for i := 0; base < bytesRead; i++ {
			blocks[i] = cipher[base : base+bestKeySize]
			base += bestKeySize
		}

		// Allocate memory for transposed matrix.
		blocksT := make([][]byte, bestKeySize)
		for j := 0; j < bestKeySize; j++ {
			blocksT[j] = make([]byte, numOfBlocks)
		}

		// Will retrive keySize chunks.
		for i, block := range blocks {
			for j := 0; j < bestKeySize; j++ {
				blocksT[j][i] = block[j]
			}
		}

		key := make([]byte, bestKeySize)

		for j := 0; j < bestKeySize; j++ {
			_, key[j] = singleByteXOR(hex.EncodeToString(blocksT[j]))
		}

		plaintext := util.RepeatingKeyXOR(string(cipher), string(key))
		score := hammingdist.PlaintextScore(plaintext)
		if score < bestScore {
			bestPlaintext = plaintext
			bestScore = score
		}
	}

	return bestPlaintext
}

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
