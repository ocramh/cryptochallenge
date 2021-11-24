/*
The file data/06_input.txt has been base64'd after being encrypted with
repeating-key XOR.

 Decrypt it.

Here's how:

Let KEYSIZE be the guessed length of the key; try values from 2 to (say) 40.
Write a function to compute the edit distance/Hamming distance between two
strings. The Hamming distance is just the number of differing bits.
The distance between:
- this is a test
and
- wokka wokka!!!
is 37. Make sure your code agrees before you proceed.

For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second KEYSIZE
worth of bytes, and find the edit distance between them. Normalize this result
by dividing by KEYSIZE.
The KEYSIZE with the smallest normalized edit distance is probably the key.
You could proceed perhaps with the smallest 2-3 KEYSIZE values. Or take 4
KEYSIZE blocks instead of 2 and average the distances.

Now that you probably know the KEYSIZE: break the ciphertext into blocks of
KEYSIZE length. Now transpose the blocks: make a block that is the first byte of
every block, and a block that is the second byte of every block, and so on.

Solve each block as if it was single-character XOR. You already have code to do
this. For each block, the single-byte XOR key that produces the best looking
histogram is the repeating-key XOR key byte for that block. Put them together
and you have the key.
*/

package set1

import (
	"errors"
	"fmt"
	"math"
)

const (
	minKeysize    = 4
	maxKeysize    = 40
	minIterations = 4
)

// HammingDistance returns the number of differing bits between a and b.
func HammingDistance(a []byte, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("input bytes buffer must be of the same length")
	}

	byteLen := 8 // number of bits in a byte

	dist := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < byteLen; j++ {
			if a[i]&(1<<j) != b[i]&(1<<j) {
				dist = dist + 1
			}
		}
	}

	return dist, nil
}

// BlockEditDistance returns the normalized hamming distance between two
// contiguous blocks of bytes taken from input. The normalized distance is
// computed by deviding the distance between the two blocks by the keysize
// itself.
func BlockEditDistance(keysize int, input []byte) (float64, error) {
	var score float64

	if keysize*minIterations > len(input) {
		return -1, fmt.Errorf("input length must be at least %d times the keysize", minIterations)
	}

	iters := (len(input) / keysize) - 1

	for i := 0; i < iters; i++ {
		block1 := input[keysize*i : keysize*(i+1)]
		block2 := input[keysize*(i+1) : keysize*(i+2)]

		dist, err := HammingDistance(block1, block2)
		if err != nil {
			return 0, err
		}

		score += float64(dist)
	}

	return score / float64(keysize) / float64(iters), nil
}

// BlockKeySize uses different key sizes to compute the edit distance of blocks
// of bytes from the inputCipher. It returns the size of the key which produced
// the smallest edit distance.
func BlockKeySize(cipher []byte) (int, error) {
	var score = float64(math.MaxInt32)
	var keysize int

	for ksize := minKeysize; ksize <= maxKeysize; ksize++ {
		distance, err := BlockEditDistance(ksize, cipher)
		if err != nil {
			return -1, err
		}

		if distance > 0 && distance < score {
			score = distance
			keysize = ksize
		}
	}

	fmt.Println(keysize)

	return keysize, nil
}

// TransposeBlocks breaks inputCipher into blocks of keysize length and
// transponses each block so that individual bytes from all blocks are grouped
// together according to their index. It returns the transposed blocks of bytes.
func TransposeBlocks(inputCipher []byte, keysize int) [][]byte {
	splitBlocks := [][]byte{}
	for i := 0; i < len(inputCipher); i += keysize {
		endIndex := i + keysize
		if endIndex > len(inputCipher) {
			endIndex = len(inputCipher)
		}

		splitBlocks = append(splitBlocks, inputCipher[i:endIndex])
	}

	transposedBlocks := make([][]byte, len(splitBlocks[0]))
	for _, block := range splitBlocks {
		for i, b := range block {
			transposedBlocks[i] = append(transposedBlocks[i], b)
		}
	}

	return transposedBlocks
}

// For each block, the single-byte XOR key that produces the best looking
// histogram is the repeating-key XOR key byte for that block. Put them together
// and you have the key.
func KeyFromXoredBlocks(blocks [][]byte) []byte {
	key := []byte{}

	for _, block := range blocks {
		var kscore float64 = 0
		var keyByte byte

		_, k, score := SingleXOR(block)
		if score > kscore {
			keyByte = k
		}

		key = append(key, keyByte)
	}

	return key
}
