package stringprobs

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// Check that a string is unique
func isUnique(str string) bool {
	dict := make(map[rune]bool)
	for _, char := range str {
		if _, exists := dict[char]; exists {
			return false
		}
		dict[char] = true
	}

	return true
}

// Given two strings check is one is a permutation of the other
func checkPermutation(a, b string) bool {
	mapA := make(map[rune]int)
	mapB := make(map[rune]int)

	for _, letter := range a {
		if _, exists := mapA[letter]; exists {
			mapA[letter]++
		} else {
			mapA[letter] = 1
		}
	}

	for _, letter := range b {
		if _, exists := mapB[letter]; exists {
			mapB[letter]++
		} else {
			mapB[letter] = 1
		}
	}
	return reflect.DeepEqual(mapA, mapB)
}

// Check if a string is a palindrome
func isPalindrome(str string) bool {
	r := regexp.MustCompile("[^A-Za-z0-9]+")
	str = strings.ToLower(r.ReplaceAllString(str, ""))

	start := 0
	end := len(str) - 1
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}

	return true
}

// One away check if a string is one edit (replace, add, or delete)
// away or none away
func oneAway(a, b string) bool {
	if len(a) != len(b) {
		return oneAddition(a, b)
	}

	edited := false
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if !edited {
				edited = true
			} else {
				return false
			}
		}
	}
	return true
}

func oneAddition(a, b string) bool {
	if (len(a)-len(b))*(len(a)-len(b)) != 1 {
		return false
	}

	var longer, shorter string
	if len(a) > len(b) {
		longer, shorter = a, b
	} else {
		longer, shorter = b, a
	}
	edited := false

	j := 0
	for i := 0; i < len(shorter); i++ {
		if longer[j] != shorter[i] {
			if !edited {
				edited = true
				i--
			} else {
				return false
			}
		}
		j++
	}
	return true
}

// Make a method that makes a string compressor
// like aaabbbc should output a3b3c
type letter struct {
	value rune
	count int
}

func stringCompress(str string) string {
	currLetter := letter{rune(str[0]), 0}
	var result []string

	for _, l := range str {
		if currLetter.value == l {
			currLetter.count++
		} else {
			result = append(result,
				fmt.Sprintf("%s%d", string(currLetter.value), currLetter.count))

			currLetter = letter{l, 1}
		}
	}
	result = append(result, fmt.Sprintf("%s%d", string(currLetter.value), currLetter.count))

	return strings.Join(result, "")
}
