package set1

import (
	"testing"
)

func TestFixedXOR(t *testing.T) {
	a := []byte("1c0111001f010100061a024b53535009181c")
	b := []byte("686974207468652062756c6c277320657965")

	expectded := "746865206b696420646f6e277420706c6179"

	got, err := XorBuffers(a, b)
	if err != nil {
		t.Fatalf("usexoedted error. %v", err)
	}

	if got != expectded {
		t.Errorf("Expected %s, got %s", expectded, got)
	}
}
