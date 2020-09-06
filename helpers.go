package helpers

import (
	"crypto/rand"
	"errors"
	"io"
)

const (
	// AlphabetUppercase an uppercase eng alphabet chars set
	AlphabetUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// AlphabetLowercase an lowercase eng alphabet chars set
	AlphabetLowercase = "abcdefghijklmnopqrstuvwxyz"
	// AlphabetDigits an alphabet of digits
	AlphabetDigits = "0123456789"
)

// GenerateRandomString generate string by length using Uppercase, Lowercase and Digits algorithms
func GenerateRandomString(length int) (str string, err error) {
	return GenerateRandomStringByAlphabet(length, AlphabetUppercase+AlphabetLowercase+AlphabetDigits)
}

// GenerateRandomStringByAlphabet generate string by certain alphabet and length
func GenerateRandomStringByAlphabet(length int, alphabet string) (str string, err error) {
	if length < 1 {
		err = errors.New("length is less than 1")
		return
	}
	chars := []byte(alphabet)
	newPword := make([]byte, length)
	randomData := make([]byte, length+(length/4)) // storage for random bytes.
	cLen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err = io.ReadFull(rand.Reader, randomData); err != nil {
			return
		}
		for _, c := range randomData {
			if c >= maxrb {
				continue
			}
			newPword[i] = chars[c%cLen]
			i++
			if i == length {
				return string(newPword), nil
			}
		}
	}
}

// RemoveStringFromSlice Remove one or more strings from []string slice and return a new array
// ...
// BenchmarkRemoveStringFromSlice20_3-4     	 2000000	       600 ns/op	     320 B/op	       1 allocs/op
// BenchmarkRemoveStringFromSlice40_11-4    	  500000	      2632 ns/op	     970 B/op	       2 allocs/op
// BenchmarkRemoveStringFromSlice40_20-4    	  500000	      3331 ns/op	    1306 B/op	       2 allocs/op
// BenchmarkRemoveStringFromSlice1000-4     	   30000	     58603 ns/op	   65568 B/op	       3 allocs/op
// BenchmarkRemoveStringFromSlice10000-4    	    2000	    600002 ns/op	  499744 B/op	       3 allocs/op
// BenchmarkRemoveStringFromSlice100000-4   	     200	   7177626 ns/op	 6897696 B/op	       3 allocs/op
//
func RemoveStringFromSlice(array []string, valuesToRemove ...string) []string {
	if len(array) == 0 || len(valuesToRemove) == 0 {
		return append([]string{}, array...)
	}
	// create a new array to save the old one
	result := make([]string, len(array), len(array))

	// factor 1.25 is used to make some slack space in map and make allocation count smaller
	hash := make(map[string]bool, int(1.25*float64(len(valuesToRemove))))
	for i := range valuesToRemove {
		hash[valuesToRemove[i]] = true
	}

	// remove elements
	j := 0
	for i := range array {
		if _, exists := hash[array[i]]; !exists {
			result[j] = array[i]
			j++
		}
	}
	result = result[0:j]
	return result
}

// FindStringInSlice find a string position in the slice
func FindStringInSlice(a string, list []string) int {
	for n, b := range list {
		if b == a {
			return n
		}
	}
	return -1
}

// StringInSlice return true if a string inside of a slice
func StringInSlice(a string, list []string) bool {
	return FindStringInSlice(a, list) >= 0
}

// Min function which returns the smaller integer from given two integers
func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// Min64 function which returns the smaller 64-bits integer from given two integers
func Min64(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

// Max function which returns the biggest integer from given two integers
func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// Max64 function which returns the biggest 64-bits integer from given two integers
func Max64(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
