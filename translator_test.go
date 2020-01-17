package main

import (
	"testing"
)

// test evaluation for string input
func TestEvaluate(t *testing.T) {
	sample := NewSampleTranslator()
	val, err := sample.evalWith("1+2")
	if err != nil {
		t.Fatal(err)
	}
	if val != "3" {
		t.Fatalf("expected 3 got %s", val)
	}
}
