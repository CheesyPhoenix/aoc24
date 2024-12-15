package day10

import "testing"

func TestExtraSimple(t *testing.T) {
	res := getResult("input_extra_simple.txt")
	if res != 1 {
		t.Fatalf("Expected 1 but got %d", res)
	}
}
func TestSimple(t *testing.T) {
	res := getResult("input_simple.txt")
	if res != 36 {
		t.Fatalf("Expected 36 but got %d", res)
	}
}

func TestSimpleP2(t *testing.T) {
	res := getResultP2("input_simple.txt")
	if res != 81 {
		t.Fatalf("Expected 81 but got %d", res)
	}
}
