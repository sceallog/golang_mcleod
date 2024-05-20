package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		log.Fatalf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(10, 5)
	if total != 5 {
		t.Errorf("Difference was incorrect, got: %d, want: %d.", total, 5)
	}
}

func TestDoMath(t *testing.T) {
	sum := doMath(5, 5, add)
	difference := doMath(10, 5, subtract)
	if sum != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 10)
	}
	if difference != 5 {
		t.Errorf("Difference was incorrect, got: %d, want: %d.", sum, 5)
	}
}
