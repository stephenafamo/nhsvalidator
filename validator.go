package nhsvalidator

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

var ErrNumberTooLarge = errors.New("number too large")

// Validate an NHS number
// If the integer has less than 10 digits, it is assumed to have leading zeros
func Validate(num uint64) error {
	if num > 9_999_999_999 {
		return ErrNumberTooLarge
	}

	var numArray [10]uint8

	// Populate the array with the digits of the number
	for i := range numArray {
		numArray[i] = uint8((num / uint64(math.Pow(10, float64(9-i)))) % 10)
	}

	return validate(numArray)
}

// An alternative version of [Validate] that takes a string
func Validate2(num string) error {
	if len(num) != 10 {
		return ErrWrongNumberOfDigits{Actual: len(num)}
	}

	// Convert the string to an integer
	numInt, err := strconv.ParseUint(num, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to convert %q to an integer: %w", num, err)
	}

	return Validate(numInt)
}

// internal validation function that takes an array of 10 integers
func validate(num [10]uint8) error {
	// If the check digit does not match the last digit of the number, it is invalid
	if getCheckDigit(num) != num[9] {
		return ErrWrongCheckDigit{Expected: getCheckDigit(num), Actual: num[9]}
	}

	return nil
}

// Ideally this would take an array of 9 integers, but since we are using
// a 10 digit number everywhere else, we use that here for convenience
func getCheckDigit(num [10]uint8) uint8 {
	var sum uint64

	for i := 0; i < 9; i++ {
		sum += uint64(num[i]) * uint64(10-i)
	}

	remainder := sum % 11
	checkDigit := 11 - remainder
	if checkDigit == 11 {
		return 0
	}

	return uint8(checkDigit)
}
