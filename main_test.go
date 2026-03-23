package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(10, 20)
	expected := 30

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

}
