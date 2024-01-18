package main

import "testing"

func TestSum(t *testing.T) {
	t.Log("Testing Sum")
	if Sum(1, 2) != 3 {
		t.Error("Expected 1 + 2 to equal 3")
	}
}

func TestSub(t *testing.T) {
	t.Log("Testing Sub")
	if Sub(2, 1) != 1 {
		t.Error("Expected 2 - 1 to equal 1")
	}
}

func TestMul(t *testing.T) {
	t.Log("Testing Mul")
	if Mul(2, 3) != 6 {
		t.Error("Expected 2 * 3 to equal 6")
	}
}

func TestDiv(t *testing.T) {
	t.Log("Testing Div")
	if Div(6, 3) != 2 {
		t.Error("Expected 6 / 3 to equal 2")
	}
}

func TestPow(t *testing.T) {
	t.Log("Testing Pow")
	result := Pow(2, 3)
	if result != 8 {
		t.Errorf("Expected 2 ^ 3 to equal 8, got %d", result)
	}
}

func TestMod(t *testing.T) {
	t.Log("Testing Mod")
	if Mod(7, 3) != 1 {
		t.Error("Expected 7 % 3 to equal 1")
	}
}
