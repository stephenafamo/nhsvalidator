package nhsvalidator

import "fmt"

type ErrWrongCheckDigit struct {
	Expected uint8
	Actual   uint8
}

func (e ErrWrongCheckDigit) Error() string {
	return fmt.Sprintf("wrong check digit: expected %d, got %d", e.Expected, e.Actual)
}

type ErrWrongNumberOfDigits struct {
	Actual int
}

func (e ErrWrongNumberOfDigits) Error() string {
	return fmt.Sprintf("wrong number of digits: expected 10, got %d", e.Actual)
}
