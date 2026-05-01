package core_validation

import "fmt"

func ValidateIsIntInBounds(number, minValue, maxValue int) error {
	if minValue <= number && number <= maxValue {
		return nil
	}
	return fmt.Errorf("number should be in bounds %d and %d, got: %d", minValue, maxValue, number)
}
