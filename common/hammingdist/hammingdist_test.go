package hammingdist

import (
	"../util"
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	result := CalculateDistance("this is a test", "wokka wokka!!!")

	if result != 37 {
		t.Errorf("Error: Expcted 37 got %v.", result)
	}
}

func TestGuessKeySize(t *testing.T) {
	// Generate Cipher
	key := "Mamta"
	plaintext := "Love you mum and dad! Life is awesome, go is amazing. This is now just filler cause I can't think of anything else to put here, I guess just blah blah blah, wow I can't write more. But I still need to keep on going oh my goodness. when will this text end - Jeez. wow that was quite a lot of text; can't believe I'm down already"

	// Test Cipher
	ciphertext := util.RepeatingKeyXOR(plaintext, key)
	guess := GuessKeySize([]byte(ciphertext))

	if guess != len(key) {
		t.Errorf("Error: expected keysize of %v, but got %v", len(key), guess)
	}
}
