// # Detect single-character XOR
//
// One of the 60-character strings in input.txt file has been encrypted by
// single-character XOR.
// Find it.

package set1

import (
	"bytes"
)

// ScoreWithSingleXOR returns src xored with key and its score based on frequency
func ScoreWithSingleXOR(src []byte, key rune) ([]byte, float64) {
	xored := bytes.Map(XorRune(key), src)
	return xored, ScoreEnglishPlaintext(string(xored))
}

// SingleXOR loops through the rune representation of alphanumeric characters,
// uses each rune to xor src and scores the xored string based on character
// frequency. It uses the rune which produced the best score to return the
// xored src.
func SingleXOR(src []byte) (xored []byte, keyChar byte, bestScore float64) {
	for i := 0; i < 256; i++ {
		key := rune(i)
		xorMapped := bytes.Map(XorRune(key), src)
		score := ScoreEnglishPlaintext(string(xorMapped))

		if score > bestScore {
			bestScore = score
			xored = xorMapped
			keyChar = byte(key)
		}
	}

	return xored, keyChar, bestScore
}
