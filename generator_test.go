package nhsvalidator

import "testing"

// Test 100 generations
// maintain a map to make sure no duplicates are generated
func TestGenerator(t *testing.T) {
	cycles := 1000
	minimum := 900
	m := make(map[int]bool)
	for i := 0; i < cycles; i++ {
		num, err := Generate()
		if err != nil {
			continue // could not generate a valid number
		}
		if m[num] {
			t.Errorf("duplicate generated: %d", num)
		}
		m[num] = true
		if err := Validate(num); err != nil {
			t.Errorf("%d. %d %v", i+1, num, err)
		}
	}

	// Make sure the generator is able to generate numbers at least 90% of the time
	if len(m) < minimum {
		t.Errorf("generator only generated %d unique numbers in %d tries", len(m), cycles)
	}
}
