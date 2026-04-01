/**
 * This test file demonstrates a TDD approach to developing software. The
 * tests use the stretchr/testify package for asserts. (Go does not have a
 * fully maintained, recognized xUnit library--Testify comes closest.)
 *
 */
package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input uint
	want  string
}

func TestInputOne(t *testing.T) {
	input := 1
	want := "one"
	got, err := IntegerToWordedString(uint(input))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, want, got)
}

func TestOneDigitInputs(t *testing.T) {
	testCases := []testCase{
		{input: 2, want: "two"},
		{input: 7, want: "seven"},
		{input: 9, want: "nine"},
	}

	for _, tc := range testCases {
		got, err := IntegerToWordedString(tc.input)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tc.want, got)
	}
}

func TestTwoDigitTeenInputs(t *testing.T) {
	testCases := []testCase{
		{input: 10, want: "ten"},
		{input: 11, want: "eleven"},
		{input: 13, want: "thirteen"},
		{input: 17, want: "seventeen"},
		{input: 19, want: "nineteen"},
	}
	for _, tc := range testCases {
		got, err := IntegerToWordedString(tc.input)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tc.want, got)
	}
}

func TestTwoDigitsGreaterThanNineteen(t *testing.T) {
	testCases := []testCase{
		{input: 21, want: "twenty-one"},
		{input: 45, want: "forty-five"},
		{input: 77, want: "seventy-seven"},
		{input: 30, want: "thirty"},
		{input: 80, want: "eighty"},
	}
	for _, tc := range testCases {
		got, err := IntegerToWordedString(tc.input)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tc.want, got)
	}
}

func TestThreeDigitNumbers(t *testing.T) {
	testCases := []testCase{
		{input: 100, want: "one-hundred"},
		{input: 300, want: "three-hundred"},
		{input: 402, want: "four-hundred two"},
		{input: 560, want: "five-hundred sixty"},
		{input: 756, want: "seven-hundred fifty-six"},
	}
	for _, tc := range testCases {
		got, err := IntegerToWordedString(tc.input)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tc.want, got)
	}
}

func TestFourDigitNumbers(t *testing.T) {
	testCases := []testCase{
		{input: 1001, want: "one-thousand, one"},
		{input: 3013, want: "three-thousand, thirteen"},
		{input: 4100, want: "four-thousand, one-hundred"},
		{input: 7068, want: "seven-thousand, sixty-eight"},
		{input: 9999, want: "nine-thousand, nine-hundred ninety-nine"},
	}
	for _, tc := range testCases {
		got, err := IntegerToWordedString(tc.input)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tc.want, got)
	}
}
