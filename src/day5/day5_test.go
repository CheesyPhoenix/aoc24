package day5

import "testing"

func TestSimple(t *testing.T) {
	res := getResult("input_simple.txt")
	if res != 143 {
		t.Fatalf("Expected 143 but got %d", res)
	}
}

func TestSimpleP2(t *testing.T) {
	res := getResultP2("input_simple.txt")
	if res != 123 {
		t.Fatalf("Expected 123 but got %d", res)
	}
}
