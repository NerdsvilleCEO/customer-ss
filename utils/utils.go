package utils

import (
	"fmt"
	sprime "github.com/otiai10/primes"
	"strings"
)

type SliceString []string
type SliceInt []int64

func (s SliceInt) contains(e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (s SliceString) contains(e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func FindVowels(input string) int {
	vowels := SliceString{"a", "e", "i", "o", "u"}
	numVals := 0
	for _, val := range input {
		if vowels.contains(strings.ToLower(fmt.Sprintf("%c", val))) {
			numVals += 1
		}
	}
	return numVals
}

func FindConsonants(input string) int {
	vowels := FindVowels(input)
	return len(input) - vowels
}

func HasCommonFactors(firstnum int64, secondnum int64) bool {
	firstfactors := SliceInt(sprime.Factorize(firstnum).All())
	secondfactors := SliceInt(sprime.Factorize(secondnum).All())
	for _, num := range firstfactors {
		if secondfactors.contains(num) {
			return true
		}
	}
	return false
}
