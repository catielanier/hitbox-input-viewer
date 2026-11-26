package controller

import "testing"

func TestEmptyAndEqual(t *testing.T) {
	a, b := Empty(), Empty()

	if !Equal(a, b) {
		t.Fatal("equal empty states should match")
	}

	a.Buttons["1p"] = true

	if Equal(a, b) {
		t.Fatal("button differences must be detected")
	}

	b.Buttons["1p"] = true

	if !Equal(a, b) {
		t.Fatal("matching states should compare equal")
	}

	a.Connected = true

	if Equal(a, b) {
		t.Fatal("connection changes must be detected")
	}
}
