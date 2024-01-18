package main

import "testing"

func TestSum(t *testing.T) {
	t.Log("Testing Sum")
	if Sum(1, 2) != 3 {
		t.Error("Expected 1 + 2 to equal 3")
	}
}
