package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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

func StringSum(input string) (output string, err error) {
	all := strings.Split(input, "")

	var result []string
	for i := 0; i < len(all); i++ {
		if all[i] == " " {
			continue
		}
		if all[i] == "+" {
			all[i] = "/"
			result = append(result, all[i])
			continue
		}
		if all[i] == "-" {
			all[i] = "/-"
		}
		result = append(result, all[i])
	}
	input = strings.Join(result, "")
	all = strings.Split((input), "/")
	result = nil
	for i := 0; i < len(all); i++ {
		if all[i] == "" {
			copy(all[i:], all[i+1:])
			all[len(all)-1] = ""
			all = all[:len(all)-1]
		}
		result = append(result, all[i])
	}
	a, err := strconv.Atoi(result[0])
	if err != nil {
		return "", fmt.Errorf("bad %w", err)
	}

	b, err := strconv.Atoi(result[1])
	if err != nil {
		return "", fmt.Errorf("bad %w", err)
	}
	if len(result) < 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	if len(result) > 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	output = fmt.Sprint(strconv.Itoa(a + b))
	return output, nil
}
