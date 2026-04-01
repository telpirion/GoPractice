/**
 * number_nlp_converters.go includes the `IntegerToWordedString` as a public
 * function.
 */

package misc

import (
	"fmt"
	"strings"
)

// IntegerToWordedString converts an unsigned integer into its English
// NLP representation
func IntegerToWordedString(number uint) (string, error) {

	// Get the digits out
	// TODO: Break this out into a new function
	thousandsDigit := number / 1000
	tmpNumber := number - (thousandsDigit * 1000)
	hundredsDigit := tmpNumber / 100
	tmpNumber = tmpNumber - (hundredsDigit * 100)
	tensDigit := tmpNumber / 10
	onesDigit := tmpNumber % 10
	tensWord := ""
	onesWord := ""

	// Handle single digit numbers
	ones := map[uint]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
		0: "zero",
	}

	onesWord = ones[onesDigit]
	if number < 10 {
		return onesWord, nil
	}

	// Handle numbers 20-99
	tensMap := map[uint]string{
		2: "twenty-",
		3: "thirty-",
		4: "forty-",
		5: "fifty-",
		6: "sixty-",
		7: "seventy-",
		8: "eighty-",
		9: "ninety-",
	}

	if tensDigit >= 2 {
		tensWord = tensMap[tensDigit]
	}

	// Handle teens
	if tensDigit == 1 {
		onesWord = ""
		tensDigit = 10 + onesDigit
		teens := map[uint]string{
			10: "ten",
			11: "eleven",
			12: "twelve",
			13: "thirteen",
			14: "fourteen",
			15: "fifteen",
			16: "sixteen",
			17: "seventeen",
			18: "eighteen",
			19: "nineteen",
		}
		tensWord = teens[tensDigit]
	}

	if onesWord == "zero" {
		onesWord = ""
		tensWord = strings.Replace(tensWord, "-", "", 1)
	}

	if number < 100 {
		fullWord := fmt.Sprintf("%s%s", tensWord, onesWord)
		return fullWord, nil
	}

	hundredsMap := map[uint]string{
		1: "one-hundred",
		2: "two-hundred",
		3: "three-hundred",
		4: "four-hundred",
		5: "five-hundred",
		6: "six-hundred",
		7: "seven-hundred",
		8: "eight-hundred",
		9: "nine-hundred",
	}

	hundredsWord := hundredsMap[hundredsDigit]

	if number < 1000 {
		fullWord := fmt.Sprintf("%s %s%s", hundredsWord, tensWord, onesWord)
		fullWord = strings.TrimRight(fullWord, " ")
		return fullWord, nil
	}

	thousandsMap := map[uint]string{
		1: "one-thousand",
		2: "two-thousand",
		3: "three-thousand",
		4: "four-thousand",
		5: "five-thousand",
		6: "six-thousand",
		7: "seven-thousand",
		8: "eight-thousand",
		9: "nine-thousand",
	}
	thousandsWord := thousandsMap[thousandsDigit]

	fullWord := fmt.Sprintf("%s, %s %s%s", thousandsWord,
		hundredsWord, tensWord, onesWord)
	fullWord = strings.TrimRight(fullWord, " ")
	fullWord = strings.Replace(fullWord, "  ", " ", -1)
	return fullWord, nil
}
