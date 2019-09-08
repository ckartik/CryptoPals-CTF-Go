package hexto64

import (
	"encoding/base64"
	"encoding/hex"
)

// hexto64 re-encodes the passed in hexString and returns it as a base64 encoding.
func HexTo64(hexString string) string {
	binaryData, err := hex.DecodeString(hexString)
	if err != nil {
		panic("There was an issue decoding the string")
	}

	return base64.StdEncoding.EncodeToString(binaryData)
}
