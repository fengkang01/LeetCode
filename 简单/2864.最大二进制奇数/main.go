package main

import "fmt"

type Test struct {
	Input  string
	Want   string
	Output string
	Passed bool
}

func main() {
	tests := []Test{
		{
			Input: "010",
			Want:  "001",
		},
		{
			Input: "0101",
			Want:  "1001",
		},
	}
	for idx, test := range tests {
		tests[idx].Output = maximumOddBinaryNumber(test.Input)
		if tests[idx].Output == test.Want {
			tests[idx].Passed = true
		}
	}
	for _, test := range tests {
		fmt.Printf("Input: %s\nWant: %s\nOutput: %s\nPassed: %v\n\n",
			test.Input, test.Want, test.Output, test.Passed)
	}
}

func maximumOddBinaryNumber(s string) string {
	count1 := 0
	newStr := ""
	for _, v := range s {
		if v == '0' {
			newStr += "0"
		} else {
			count1++
		}
	}
	for count1 > 1 {
		newStr = "1" + newStr
		count1--
	}
	return newStr + "1"
}

// 要求是奇数所以最低位是1，如果字符串种包含多个1则除了第一个1剩余全部放在最前边
