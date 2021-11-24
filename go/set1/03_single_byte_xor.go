// # Challenge 3: Single-byte XOR cipher
//
// The hex encoded string:
// 1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
// ... has been XOR'd against a single character. Find the key, decrypt the message.
// You can do this by hand. But don't: write code to do it for you.
// How? Devise some method for "scoring" a piece of English plaintext. Character
// frequency is a good metric. Evaluate each output and choose the one with the best score.
//
// ## Tips
// For the scoring algorithms see the [Letter_frequency](https://en.wikipedia.org/wiki/Letter_frequency) entry on Wikipedia.

package set1

import (
	"encoding/hex"
	"unicode"
)

var (
	charValues = map[rune]float64{
		'A': 7.8,
		'B': 2,
		'C': 4,
		'D': 3.8,
		'E': 11,
		'F': 1.4,
		'G': 3,
		'H': 2.3,
		'I': 8.6,
		'J': 0.21,
		'K': 0.97,
		'L': 5.3,
		'M': 2.7,
		'N': 7.2,
		'O': 6.1,
		'P': 2.8,
		'Q': 0.19,
		'R': 7.3,
		'S': 8.7,
		'T': 6.7,
		'U': 3.3,
		'V': 1,
		'W': 0.91,
		'X': 0.27,
		'Y': 1.6,
		'Z': 0.44,
		' ': 12.0,
	}
)

func ScoreEnglishPlaintext(i string) float64 {
	var score float64 = 0
	for _, c := range i {
		val, ok := charValues[unicode.ToUpper(c)]
		if ok {
			score = score + val
		}
	}

	return score
}

// DecodeHexStr decodes a hex encoded string into a slice of bytes.
func DecodeHexStr(src string) ([]byte, error) {
	dest := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dest, []byte(src))
	if err != nil {
		return nil, err
	}

	return dest, nil
}

// XorRune is a mapper function that returning a runction which xors the input
// with k.
// Note: xor is both commutative (a × b = b × a) and
// associative ( a × b ) × c = a × ( b × c )
func XorRune(key rune) func(rune) rune {
	return func(char rune) rune {
		return char ^ key
	}
}

// XorBytes xor each byte in src using key.
func XorBytes(src []byte, key byte) string {
	var output = make([]byte, len(src))
	for i := 0; i < len(src); i++ {
		xored := src[i] ^ byte(key)
		output[i] = xored
	}

	return string(output)
}
