package day6

import "testing"

func TestSimple(t *testing.T) {
	res := getResult("input_simple.txt")
	if res != 41 {
		t.Fatalf("Expected 41 but got %d", res)
	}
}

func TestSimpleP2(t *testing.T) {
	res := getResultP2("input_simple.txt")
	if res != 6 {
		t.Fatalf("Expected 6 but got %d", res)
	}
}
