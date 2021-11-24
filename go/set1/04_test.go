package set1

import (
	"bufio"
	"os"
	"testing"
)

const (
	inputFilePath = "data/04_input.txt"
)

func Test04(t *testing.T) {
	f2, err := os.Open(inputFilePath)
	if err != nil {
		t.Fatal(err)
	}
	defer f2.Close()

	scanner := bufio.NewScanner(f2)
	var bestScore float64 = 0
	var decoded string = ""

	for scanner.Scan() {
		fromHex := DecodeHexToBytes(scanner.Text())

		xored, _, score := SingleXOR(fromHex)
		if score > bestScore {
			bestScore = score
			decoded = string(xored)
		}
	}

	t.Logf("decoded: %s", decoded)
}
