package main

import (
	"../../common/attack"
	"../../common/hammingdist"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"os"
)

// TODO: Break up into Returning Keysize function and creating blocks, only pass a slice of the cipher into this and all other funcitions.
func breakRepeatingKeyXOR() {
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

	// Init paramters for discovery.
	bestResult := math.Inf(1)
	bestKeySize := 5

	// Discover value for the keysize.
	// TODO: Fix this as it's not finding the right key size.
	for keySize := 2; keySize < 41; keySize++ {
		distanceMeasure := 0.0
		base := 0
		chunk1 := string(cipher[base : base+keySize])
		base += keySize
		chunk2 := string(cipher[base : base+keySize])
		base += keySize
		chunk3 := string(cipher[base : base+keySize])
		base += keySize
		chunk4 := string(cipher[base : base+keySize])

		distanceMeasure += hammingdist.CalculateDistance(chunk1, chunk2)
		distanceMeasure += hammingdist.CalculateDistance(chunk1, chunk3)
		distanceMeasure += hammingdist.CalculateDistance(chunk1, chunk4)
		distanceMeasure += hammingdist.CalculateDistance(chunk2, chunk3)
		distanceMeasure += hammingdist.CalculateDistance(chunk2, chunk4)
		distanceMeasure += hammingdist.CalculateDistance(chunk3, chunk4)

		result := float64(distanceMeasure / (float64(keySize) * 6.0))
		if result < bestResult {
			bestKeySize = keySize
			bestResult = result
		}
	}
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
		_, key[j] = attack.SingleByteXOR(hex.EncodeToString(blocksT[j]))
	}

	fmt.Println(decryptRepeatingKeyXOR(cipher, key))

}

func decryptRepeatingKeyXOR(byteStream, key []byte) string {
	cipherStream := make([]byte, len(byteStream))
	keySize := len(key)
	for i := range byteStream {
		cipherStream[i] = byteStream[i] ^ key[i%keySize]
	}

	return string(cipherStream)
}

func main() {
	breakRepeatingKeyXOR()
}
