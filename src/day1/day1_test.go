package day1

import "testing"

func TestSimple(t *testing.T) {
	res := getResult("input_simple.txt")
	if res != 11 {
		t.Fatalf("Expected 11 but got %d", res)
	}
}

func TestSimpleP2(t *testing.T) {
	res := getResultP2("input_simple.txt")
	if res != 31 {
		t.Fatalf("Expected 31 but got %d", res)
	}
}
