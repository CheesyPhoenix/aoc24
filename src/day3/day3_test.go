package day3

import "testing"

func TestSimple(t *testing.T) {
	res := getResult("input_simple.txt")
	if res != 161 {
		t.Fatalf("Expected 161 but got %d", res)
	}
}

func TestSimpleP2(t *testing.T) {
	res := getResultP2("input_simple2.txt")
	if res != 48 {
		t.Fatalf("Expected 48 but got %d", res)
	}
}
