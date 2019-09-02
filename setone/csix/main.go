package main

import (
	"../../common/hammingdist"
	"fmt"
	"os"
)

func breakRepeatingKeyXOR() {
	file, err := os.Open("6.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	cipherChunk := make([]byte, 160)
	file.Read(cipherChunk)

	bestResult := 5.0
	bestKey := 3

	for keySize := 2; keySize < 41; keySize++ {
		distanceMeasure := 0.0
		base := 0
		chunk1 := string(cipherChunk[base : base+keySize])
		base += keySize
		chunk2 := string(cipherChunk[base : base+keySize])
		base += keySize
		chunk3 := string(cipherChunk[base : base+keySize])
		base += keySize
		chunk4 := string(cipherChunk[base : base+keySize])

		distanceMeasure += hammingdist.CalculateDistance(chunk1, chunk2)
		distanceMeasure += hammingdist.CalculateDistance(chunk1, chunk3)
		distanceMeasure += hammingdist.CalculateDistance(chunk1, chunk4)
		distanceMeasure += hammingdist.CalculateDistance(chunk2, chunk3)
		distanceMeasure += hammingdist.CalculateDistance(chunk2, chunk4)
		distanceMeasure += hammingdist.CalculateDistance(chunk3, chunk4)

		result := float64(distanceMeasure / (float64(keySize) * 6.0))
		if result < bestResult {
			bestKey = keySize
			bestResult = result
		}
	}
	fmt.Println(bestKey)
}

func main() {
	breakRepeatingKeyXOR()
}
