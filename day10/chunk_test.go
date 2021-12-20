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
		{"[}", "}]", errLineCorrupted},
		{"<>>", ">", nil}, // single chunk ok
		{"<[]>", "", nil},
		{"", "", nil},
		{">", ">", errLineCorrupted},
		{"{", "}", errLineIncomplete},
		{"(<({[]})>)", "", nil},
		{"(<({[>]})>)", ">]})>)", errLineCorrupted},
		{"(<{<[()", "]>}>)", errLineIncomplete},
		{"(<{<[[]", "]>}>)", errLineIncomplete},
		{"(<{<[{}", "]>}>)", errLineIncomplete},
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

func TestScore(t *testing.T) {

	for _, test := range []struct {
		input string
		score int
	}{
		{"}}]])})]", 288957},
		{")}>]})", 5566},
		{"}}>}>))))", 1480781},
		{"]]}}]}]}>", 995444},
		{"])}>", 294},
	} {
		result := score(test.input)
		if result != test.score {
			t.Errorf("scoring %s, expected %d, got %d", test.input, test.score, result)
		}
	}
}
