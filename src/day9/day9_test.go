package day9

import "testing"

func TestSimple(t *testing.T) {
	res := getResult("input_simple.txt")
	if res != 1928 {
		t.Fatalf("Expected 1928 but got %d", res)
	}
}

func TestSimpleP2(t *testing.T) {
	res := getResultP2("input_simple.txt")
	if res != 2858 {
		t.Fatalf("Expected 2858 but got %d", res)
	}
}
