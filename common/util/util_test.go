package util

import "testing"

func TestHexTo64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if output != HexTo64(input) {
		t.Errorf("Error: There was an issue converting the HexTo64") // to indicate test failed.
	}
}

func TestRepeatingKeyXOR(t *testing.T) {
	S1C5Input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	S1C5Answer := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	S1C5Result := RepeatingKeyXOR(S1C5Input, "ICE")
	if S1C5Result != S1C5Answer {
		t.Errorf("Error: Expected:\n%v, but received:\n%v", S1C5Answer, S1C5Result)
	}

}
