package day4

import "testing"

func TestSimple(t *testing.T) {
	res := getResult("input_simple.txt")
	if res != 18 {
		t.Fatalf("Expected 18 but got %d", res)
	}
}
func TestExtraSimple(t *testing.T) {
	res := getResult("input_extra_simple.txt")
	if res != 4 {
		t.Fatalf("Expected 4 but got %d", res)
	}
}

func TestSimpleP2(t *testing.T) {
	res := getResultP2("input_simple.txt")
	if res != 9 {
		t.Fatalf("Expected 9 but got %d", res)
	}
}
