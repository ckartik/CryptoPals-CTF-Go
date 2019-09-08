package hexto64

import "testing"

func TestHexTo64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if output != HexTo64(input) {
		t.Errorf("Error: There was an issue converting the HexTo64") // to indicate test failed.
	}
}
