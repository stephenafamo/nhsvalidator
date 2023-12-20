package nhsvalidator

import (
	_ "embed"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

//go:embed valid.txt
var valid string

func TestValidator(t *testing.T) {
	for lineNum, num := range strings.Split(valid, "\n") {
		if num == "" || strings.HasPrefix(num, "#") {
			continue
		}

		// Check if it is a valid number
		if err := Validate2(num); err != nil {
			t.Errorf("%d. %q %v", lineNum+1, num, err)
		}

		// Now we check if we get an error when we change
		// the last digit to any other number

		// convert the last digit to an integer
		correct, err := strconv.Atoi(string(num[9]))
		if err != nil {
			panic(err) // this should never happen
		}

		var expErr ErrWrongCheckDigit
		// Cycle through all the other invalid digits
		for i := 1; i <= 9; i++ {
			newLastDigit := (correct + i) % 10
			expectedErr := ErrWrongCheckDigit{
				Expected: correct,
				Actual:   newLastDigit,
			}

			num2 := num[:9] + strconv.Itoa(newLastDigit)
			err := Validate2(num2)

			if !errors.As(err, &expErr) ||
				expErr.Expected != correct ||
				expErr.Actual != newLastDigit {
				t.Errorf(
					"%q should have wrong check digit\nexpected: %v\ngot: %v",
					num2, expectedErr, err,
				)
			}
		}
	}
}

func TestLargeNumber(t *testing.T) {
	for i := 0; i < 100; i++ {
		num := rand.Int() + 9999999999
		err := Validate(num)
		if !errors.Is(err, ErrNumberTooLarge) {
			t.Errorf("%d should not be valid but passed validation", num)
		}
	}
}

// TestWrongNumberOfDigits tests strings that are not 10 digits long
func TestWrongNumberOfDigits(t *testing.T) {
	for i := 1; i <= 20; i++ {
		if i == 10 {
			continue
		}
		s := strings.Repeat("1", i)
		err := Validate2(s)
		if !errors.Is(err, ErrWrongNumberOfDigits{Actual: i}) {
			t.Errorf("%q should not be valid but passed validation", s)
		}
	}
}

// TestBadString tests strings that are not valid integers
func TestBadString(t *testing.T) {
	for _, s := range []string{
		"123456789a",
		"123a567890",
	} {
		err := Validate2(s)
		if err == nil {
			t.Errorf("%q should not be valid but passed validation", s)
		}
	}
}
