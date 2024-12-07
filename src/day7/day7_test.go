package day7

import "testing"

func TestSimple(t *testing.T) {
	res := getResult("input_simple.txt")
	if res != 3749 {
		t.Fatalf("Expected 3749 but got %d", res)
	}
}

func TestSimpleP2(t *testing.T) {
	res := getResultP2("input_simple.txt")
	if res != 11387 {
		t.Fatalf("Expected 11387 but got %d", res)
	}
}
