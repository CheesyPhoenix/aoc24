package day8

import "testing"

func TestSimple(t *testing.T) {
	res := getResult("input_simple.txt")
	if res != 14 {
		t.Fatalf("Expected 14 but got %d", res)
	}
}

func TestSimpleP2(t *testing.T) {
	res := getResultP2("input_simple.txt")
	if res != 34 {
		t.Fatalf("Expected 34 but got %d", res)
	}
}
