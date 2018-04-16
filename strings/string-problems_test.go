package stringprobs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var stringTestCases = []struct {
	word     string
	expected bool
}{
	{
		word:     "computer",
		expected: true,
	},
	{
		word:     "brown cat",
		expected: true,
	},
	{
		word:     "hello there!",
		expected: false,
	},
	{
		word:     "I'm ok ",
		expected: false,
	},
}

func TestIsUnique(t *testing.T) {
	for _, test := range stringTestCases {
		outcome := isUnique(test.word)
		assert.Equal(t, test.expected, outcome)
	}
}

var permutationTests = []struct {
	first    string
	second   string
	expected bool
}{
	{
		first:    "rat",
		second:   "tar",
		expected: true,
	},
	{
		first:    "cartoon",
		second:   "tooncar",
		expected: true,
	},
	{
		first:    "cats",
		second:   "cat",
		expected: false,
	},
}

func TestCheckPermutation(t *testing.T) {
	for _, test := range permutationTests {
		actual := checkPermutation(test.first, test.second)
		assert.Equal(t, test.expected, actual)
	}
}

var testPalindrome = []struct {
	input    string
	expected bool
}{
	{
		input:    "A car, a man, a maraca",
		expected: true,
	},
	{
		input:    "Hello World",
		expected: false,
	},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range testPalindrome {
		actual := isPalindrome(test.input)
		assert.Equal(t, test.expected, actual)
	}
}

var testCasesOneAway = []struct {
	a        string
	b        string
	expected bool
}{
	{
		a:        "cat",
		b:        "pat",
		expected: true,
	},
	{
		a:        "cartoon",
		b:        "cart",
		expected: false,
	},
	{
		a:        "per",
		b:        "peer",
		expected: true,
	},
	{
		a:        "pr",
		b:        "peer",
		expected: false,
	},
	{
		a:        "mister",
		b:        "mr",
		expected: false,
	},
}

func TestOneAway(t *testing.T) {
	for _, test := range testCasesOneAway {
		actual := oneAway(test.a, test.b)
		assert.Equal(t, test.expected, actual)
	}
}

var testStringCompress = []struct {
	input    string
	expected string
}{
	{
		input:    "hellllno",
		expected: "h1e1l4n1o1",
	},
}

func TestStringCompress(t *testing.T) {
	for _, test := range testStringCompress {
		actual := stringCompress(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
