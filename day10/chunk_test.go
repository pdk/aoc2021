package main

import (
	"testing"
)

func TestParse(t *testing.T) {

	for _, testCase := range []struct {
		input, result string
		err           error
	}{
		{"[]", "", nil},
		{"[}", "}", errLineCorrupted},
		{"<>>", ">", nil}, // single chunk ok
		{"<[]>", "", nil},
		{"", "", nil},
		{">", ">", errLineCorrupted},
		{"{", "", errLineIncomplete},
		{"(<({[]})>)", "", nil},
		{"(<({[>]})>)", ">", errLineCorrupted},
	} {
		result, err := parse(testCase.input)
		if err != testCase.err {
			t.Errorf("parsing %#v, expected err %#v, got %#v", testCase.input, testCase.err, err)
		}
		if result != testCase.result {
			t.Errorf("parsing %#v, expected result %#v, got %#v", testCase.input, testCase.result, result)
		}
	}
}
