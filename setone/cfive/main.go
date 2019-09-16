package main

import (
	"encoding/hex"
	"fmt"
)

func repeatingKeyXOR(plaintext, key string) string {
	byteStream := []byte(plaintext)
	cipherStream := make([]byte, len(byteStream))
	byteKey := []byte(key)
	keySize := len(byteKey)
	for i := range byteStream {
		cipherStream[i] = byteStream[i] ^ byteKey[i%keySize]
	}

	return hex.EncodeToString(cipherStream)
}

func main() {
	S1C5Input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	S1C5Answer := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	S1C5Result := repeatingKeyXOR(S1C5Input, "ICE")
	if S1C5Result != S1C5Answer {
		fmt.Printf(S1C5Result)
	}

}
