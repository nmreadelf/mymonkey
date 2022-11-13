package object

import "testing"

func TestStringHashKey(t *testing.T) {
	h1 := &String{"Hello World"}
	h2 := &String{"Hello World"}
	diff1 := String{"My name is johnny"}
	diff2 := String{"My name is johnny"}

	if h1.HashKey() != h2.HashKey() {
		t.Errorf("strings with same content have differen hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have differen hash keys")
	}

	if h1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
