package nhsvalidator

import (
	"errors"
	"math"
	"math/rand"
)

func Generate() (uint64, error) {
	// 10 tries to generate a valid number
	for i := 0; i < 10; i++ {
		// The array of 10 integers that make up the number
		// This will be sent to getCheckDigit to get the check digit
		var nums [10]uint8

		for i := 0; i < 10; i++ {
			nums[i] = uint8(rand.Intn(10))
		}

		// Add the check digit
		nums[9] = getCheckDigit(nums)

		if nums[9] == 10 {
			// If the check digit is 10, the number is invalid
			// so we try again
			continue
		}

		// Convert the array of integers to a single integer
		var num uint64 = 0
		for i, n := range nums {
			num += uint64(n) * uint64(math.Pow(10, float64(9-i)))
		}

		return num, nil
	}

	return 0, errors.New("failed to generate a valid number after 10 tries")
}
