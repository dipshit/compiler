package main

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	sample := &SampleTranslator{}
	val, err := sample.evaluate("1+2")
	if err != nil {
		t.Fatal(err)
	}
	if val != "3" {
		t.Fatalf("expected 3 got %s", val)
	}
}
