package main

import (
	"fmt"
	"strings"
)

type Test struct {
	Inputer string
	Wanter  string
	Getter  string
	Passed  bool
}

func main() {
	tests := []Test{
		{
			Inputer: "capiTalIze tHe titLe",
			Wanter:  "Capitalize The Title",
		},
		{
			Inputer: "First leTTeR of EACH Word",
			Wanter:  "First Letter of Each Word",
		},
		{
			Inputer: "i lOve leetcode",
			Wanter:  "i Love Leetcode",
		},
	}
	for idx, test := range tests {
		got := capitalizeTitle(test.Inputer)
		tests[idx].Getter = got
		tests[idx].Passed = got == test.Wanter
	}
	for _, test := range tests {
		fmt.Printf("Input: %s\nOutput: %s\nGet: %s\nPassed: %v\n\n",
			test.Inputer, test.Wanter, test.Getter, test.Passed)
	}
}

func capitalizeTitle(title string) string {
	words := strings.Split(title, " ")
	for i, word := range words {
		words[i] = help(word)
	}
	return strings.Join(words, " ")
}

func help(word string) string {
	if len(word) <= 2 {
		return strings.ToLower(word)
	}
	return strings.ToUpper(word[0:1]) + strings.ToLower(word[1:])
}
