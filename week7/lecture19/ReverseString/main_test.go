package main

import (
	"bytes"
	"io"
	"testing"
)

func TestCheckReversing(t *testing.T) {
	input := "Hello world!"
	expectedResult := "!dlrow olleH"

	ourReverseStringReader := NewReverseStringReader(input)
	var buffer bytes.Buffer
	io.Copy(&buffer, ourReverseStringReader)

	result := buffer.String()

	if len(expectedResult) != len(result) {
		t.Errorf("Expected result: %s", expectedResult)
	} else {
		for i := 0; i < len(expectedResult); i++ {
			if expectedResult[i] != result[i] {
				t.Errorf("Expected element: %s, ", string(expectedResult[i]))
			}
		}
	}
}
