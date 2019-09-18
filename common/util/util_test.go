package util

import "testing"

func TestRepeatingKeyXOR(t *testing.T) {
	S1C5Input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	S1C5Answer := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	S1C5Result := RepeatingKeyXOR(S1C5Input, "ICE")
	if S1C5Result != S1C5Answer {
		t.Errorf("Error: Expected:\n%v, but received:\n%v", S1C5Answer, S1C5Result)
	}

}
