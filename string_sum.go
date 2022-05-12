package string_sum

import (
	"errors"
	"fmt"
	sc "strconv"
	sts "strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (string, error) {
	fmt.Println("--------------------------")
	input = sts.ReplaceAll(input, " ", "")
	if len(input) == 0 {
		return "", errorEmptyInput
	}
	var sum = 0
	fmt.Println(input)
	input = input[0:1] + sts.ReplaceAll(input[1:], "-", "+-")
	fmt.Println(input)
	var split = sts.Split(input, "+")
	fmt.Println(split)
	if len(split) == 2 {
		if arg1, e1 := sc.Atoi(split[0]); e1 == nil {
			sum += arg1
		} else {
			return "", fmt.Errorf("invalid input: %v", e1)
		}
		if arg2, e2 := sc.Atoi(split[1]); e2 == nil {
			sum += arg2
		} else {
			return "", fmt.Errorf("invalid input: %v", e2)
		}
	} else {
		return "", errorNotTwoOperands
	}
	return sc.Itoa(sum), nil
}
