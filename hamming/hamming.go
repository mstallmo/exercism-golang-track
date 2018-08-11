//The hamming package has a function Distance that takes two even length strings and compares them for hamming distance
package hamming

import (
	"errors"
)

//Distance takes in two strings of equal length and returns the hamming distance. If the strings are of unequal length
//Distance will throw an error otherwise if successful it will return an int of the hamming distance value
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("sequence lengths are not equal")
	}

	diff := 0
	for i := 0; i < len(a); i++ {
		if string(a[i]) != string(b[i]) {
			diff += 1
		}
	}
	return diff, nil
}
