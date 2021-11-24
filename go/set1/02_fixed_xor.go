// # Challenge 2: Fixed XOR
//
// Write a function that takes two equal-length buffers and produces their XOR
// combination. If your function works properly, then when you feed it the
// string:
// 1c0111001f010100061a024b53535009181c
// ... after hex decoding, and when XOR'd against:
// 686974207468652062756c6c277320657965
// ... should produce:
// 746865206b696420646f6e277420706c6179

package set1

import (
	"encoding/hex"
	"errors"
)

func XorBuffers(a []byte, b []byte) (string, error) {
	if len(a) != len(b) {
		return "", errors.New("input buffers must be equal in size")
	}

	aBuf := make([]byte, hex.DecodedLen(len(a)))
	_, err := hex.Decode(aBuf, a)
	if err != nil {
		return "", err
	}

	bBuf := make([]byte, hex.DecodedLen(len(b)))
	_, err = hex.Decode(bBuf, b)
	if err != nil {
		return "", err
	}

	output := make([]byte, len(aBuf))
	for i := 0; i < len(aBuf); i++ {
		output[i] = aBuf[i] ^ bBuf[i]
	}

	return hex.EncodeToString(output), nil
}
