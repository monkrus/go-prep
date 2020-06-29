package main

import (
	"testing"
)

func TestInitFunc(t *testing.T) {
	input := "init function testing"
	expectedOutput := "init_function_testing"

	actual, err := init(input)
	if err != nil {
		t.Errorf("expected no error, bit gpt %v", err)

	}
	if actual != expectedOutput {
		t.Errorf("expected output to be %s, but got %s", expectedOutput, actual)
	}
}
