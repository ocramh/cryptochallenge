package set1

import (
	"bytes"
	"testing"
)

const (
	hexInput_03 = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
)

func TestSingleByteXORCipher(t *testing.T) {
	b, err := DecodeHexStr(hexInput_03)
	if err != nil {
		panic(err)
	}

	var decoded string
	var key byte
	var decodedScore float64

	for i := 0; i < 256; i++ {
		r := rune(i)

		xored := bytes.Map(XorRune(r), b)
		score := ScoreEnglishPlaintext(string(xored))
		if score > decodedScore {
			decodedScore = score
			key = byte(r)
			decoded = string(xored)
		}
	}

	t.Logf("decoded %s. key %s", decoded, string(key))
}
