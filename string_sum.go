package string_sum

import (
	"errors"
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
		input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")

	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	all := strings.Split(input, "")
	if len(all) < 3 || len(all) > 4 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	first := 0
	second := 0
	ans := 0
	if v, err := strconv.Atoi(all[0]); err == nil {
		first += v
		if all[1] == "+" {
			if v, err := strconv.Atoi(all[2]); err == nil {
				ans = first + v
				second += v

			} else {
				fmt.Println("error")
			}
		} else if all[1] == "-" {
			if v, err := strconv.Atoi(all[2]); err == nil {
				ans = first - v
				second -= v
			} else {
				fmt.Println("error")
			}
		}
	} else if all[0] == "-" {
		if v, err := strconv.Atoi(all[1]); err == nil {
			first -= v
			if all[1] == "+" {
				if v, err := strconv.Atoi(all[2]); err == nil {
					ans = first + v
					second += v
				} else {
					return "", fmt.Errorf("bad %w", err)
				}
			} else if all[2] == "-" {
				if v, err := strconv.Atoi(all[3]); err == nil {
					ans = first - v
					second -= v
				} else {
					return "", fmt.Errorf("bad %w", err)
				}
			}
		} else {
			return "", fmt.Errorf("bad %w", err)
		}
	} else {
		return "", fmt.Errorf("bad %w", err)
	}
	
	return fmt.Sprint(strconv.Itoa(ans)), nil
}
