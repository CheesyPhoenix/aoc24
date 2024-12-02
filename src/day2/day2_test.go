package day2

import "testing"

func TestSimple(t *testing.T) {
	res := getResult("input_simple.txt")
	if res != 2 {
		t.Fatalf("Expected 2 but got %d", res)
	}
}

func TestSimpleP2(t *testing.T) {
	res := getResultP2("input_simple.txt")
	if res != 6 {
		t.Fatalf("Expected 6 but got %d", res)
	}
}
