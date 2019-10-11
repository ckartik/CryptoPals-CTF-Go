package util

import (
	"bufio"
	b64 "encoding/base64"
	"io"
	"os"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	S1C2Input := "1c0111001f010100061a024b53535009181c"
	S1C2Input2 := "686974207468652062756c6c277320657965"
	S1C2Answer := "746865206b696420646f6e277420706c6179"
	S1C2Result := FixedXOR(S1C2Input, S1C2Input2)
	if S1C2Result != S1C2Answer {
		t.Errorf("Error: Expected 746865206b696420646f6e277420706c6179, got %v", S1C2Result)
	}
}

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

// TODO: Unable to decode full-file into buffer.
func TestDecryptAES(t *testing.T) {
	fh, err := os.Open("./7.txt")
	if err != nil {
		t.Errorf("Test failed because file could not be opened.\n%v", err)
	}
	defer fh.Close()

	stats, err := fh.Stat()
	if err != nil {
		t.Errorf("Test failed because file stats could not be opened.\n%v", err)
	}
	reader := b64.NewDecoder(b64.StdEncoding, fh)
	size := stats.Size()
	t.Logf("File has a size of %v bytes.", size)

	cipher := make([]byte, size)
	bytesRead, err := io.ReadFull(reader, cipher)
	t.Logf("%v bytes read into buffer after decoding base64.", bytesRead)

	t.Logf("Value returned:\n%v", DecryptAES(cipher[:bytesRead], []byte("YELLOW SUBMARINE")))

}

func TestPKCS7(t *testing.T) {
	testStr1 := "Hello World"
	testStr2 := "Hello WorldWorld"
	testStr3 := "Hello WorldWorldWorld"
	testStr4 := "Hello WorldWorldWorldWorld"

	if len(PKCS7([]byte(testStr1), BlockSize16))%BlockSize16 != 0 {
		t.Error("Padding was not sucessful")
	}

	if len(PKCS7([]byte(testStr2), BlockSize16))%BlockSize16 != 0 {
		t.Error("Padding was not sucessful")
	}

	if len(PKCS7([]byte(testStr3), BlockSize16))%BlockSize16 != 0 {
		t.Error("Padding was not sucessful")
	}

	if len(PKCS7([]byte(testStr4), BlockSize16))%BlockSize16 != 0 {
		t.Error("Padding was not sucessful")
	}
	if len(PKCS7([]byte(testStr1), BlockSize16))%BlockSize16 != 0 {
		t.Error("Padding was not sucessful")
	}

	if len(PKCS7([]byte(testStr2), BlockSize8))%BlockSize8 != 0 {
		t.Error("Padding was not sucessful")
	}

	if len(PKCS7([]byte(testStr3), BlockSize8))%BlockSize8 != 0 {
		t.Error("Padding was not sucessful")
	}

	if len(PKCS7([]byte(testStr4), BlockSize8))%BlockSize8 != 0 {
		t.Error("Padding was not sucessful")
	}

}

func TestDetectAES(t *testing.T) {
	fh, err := os.Open("./8.txt")
	defer fh.Close()
	if err != nil {
		t.Errorf("Test failed because file could not be opened.\n%v", err)
	}

	// TODO: Fix this to use different conncurrent architecture, this will cause a deadlock if the buffer is not large enough.
	ciphers := make(chan []byte, 300)
	detected := DetectAES(ciphers)

	scanner := bufio.NewScanner(fh)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ciphers <- []byte(scanner.Text())
	}

	close(ciphers)
	stringFound := <-detected
	t.Log(stringFound)
}
