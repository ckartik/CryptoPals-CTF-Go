package attack

import (
	"os"
	"testing"
)

func TestDetectSingleKeyXOR(t *testing.T) {
	file, err := os.Open("4.txt")
	if err != nil {
		panic(err)
	}
	S1C4Result := DetectSingleKeyXOR(file)
	S1C4Answer := "Now that the party is jumping\n"
	if S1C4Result != S1C4Answer {
		t.Errorf("Error: expected %v, but recieved %v", S1C4Answer, S1C4Result)
	}
}
