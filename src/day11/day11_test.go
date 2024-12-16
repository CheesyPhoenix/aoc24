package day11

import "testing"

func TestSimple(t *testing.T) {
	res := getResult("input_simple.txt")
	if res != 55312 {
		t.Fatalf("Expected 55312 but got %d", res)
	}
}

// func TestSimpleP2(t *testing.T) {
// 	res := getResultP2("input_simple.txt")
// 	if res != 55312 {
// 		t.Fatalf("Expected 55312 but got %d", res)
// 	}
// }
