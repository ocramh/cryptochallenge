package set1

import (
	"encoding/base64"
	"testing"
)

const (
	hexInput_01 = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
)

func TestConvertHexToBase64(t *testing.T) {

	decodedHex := DecodeHexToBytes(hexInput_01)
	b64Encoded := base64.StdEncoding.EncodeToString(decodedHex)

	expectedB64Str := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	if b64Encoded != expectedB64Str {
		t.Errorf("Expected %s, got %s", expectedB64Str, b64Encoded)
	}
}
