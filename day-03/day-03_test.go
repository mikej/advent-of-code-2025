package main

import "testing"

func TestMaxJoltageWhereDigitsAreTogether(t *testing.T) {
	checkMaxJoltageForPart1(t, "987654321111111", 98)
}

func TestMaxJoltageWhereDigitsAreAtEachEnd(t *testing.T) {
	checkMaxJoltageForPart1(t, "811111111111119", 89)
}

func TestMaxJoltageWhereDigitsAreSeparatedPartWayThrough(t *testing.T) {
	checkMaxJoltageForPart1(t, "818181911112111", 92)
}

func TestMaxJoltageReturnsAnErrorWhenStringIsTooShort(t *testing.T) {
	result, err := maxJoltage2("1", 2)
	if err == nil {
		t.Errorf("expected to get an error, got no error and a result of %d", result)
	}
}

func TestMaxJoltageForPart2LeavingSome1sAtTheEnd(t *testing.T) {
	checkMaxJoltageForPart2(t, "987654321111111", 987654321111)
}

func TestMaxJoltageForpart2LeavingSome1sInTheMiddle(t *testing.T) {
	checkMaxJoltageForPart2(t, "234234234234278", 434234234278)
}

func TestMaxJoltageForPart2LeavingSomeLowDigits(t *testing.T) {
	checkMaxJoltageForPart2(t, "234234234234278", 434234234278)
}

func TestMaxJoltageForPart2LeavingSome1sNearTheFront(t *testing.T) {
	checkMaxJoltageForPart2(t, "818181911112111", 888911112111)
}

func checkMaxJoltageForPart1(t *testing.T, str string, expected int) {
	result, err := maxJoltage2(str, 2)
	if err != nil {
		t.Errorf("unexpected error getting max joltage: %v", err)
	}
	if result != expected {
		t.Errorf("max joltage should be %d, got %d", expected, result)
	}
}

func checkMaxJoltageForPart2(t *testing.T, str string, expected int) {
	result, err := maxJoltage2(str, 12)
	if err != nil {
		t.Errorf("unexpected error getting max joltage: %v", err)
	}
	if result != expected {
		t.Errorf("max joltage should be %d, got %d", expected, result)
	}
}
