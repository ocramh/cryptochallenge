package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestHamminDistance(t *testing.T) {
	a := []byte(`this is a test`)
	b := []byte(`wokka wokka!!!`)
	expectedScore := 37

	got, err := HammingDistance(a, b)
	if err != nil {
		t.Error("unexpected error", err)
	}

	if got != expectedScore {
		t.Fail()
	}
}

func Test06(t *testing.T) {
	f, err := os.Open("data/06_input.txt")
	if err != nil {
		t.Error(err)
	}

	b64Decoded := base64.NewDecoder(base64.StdEncoding, f)
	cipher, err := ioutil.ReadAll(b64Decoded)
	if err != nil {
		t.Error(err)
	}

	keySize, err := BlockKeySize(cipher)
	if err != nil {
		t.Error(err)
	}

	key := KeyFromXoredBlocks(TransposeBlocks(cipher, keySize))

	plaintext, err := hex.DecodeString(XorWithRepeatingKey(cipher, key))
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(plaintext))
}
