package main

import "testing"

func TestRepeatedTwice(t *testing.T) {
	shouldBeInvalid(12341234, t)
}

func TestRepeatedThreeTimes(t *testing.T) {
	shouldBeInvalid(123123123, t)
}

func TestRepeatedFiveTimes(t *testing.T) {
	shouldBeInvalid(1212121212, t)
}

func TestRepeatedSevenTimes(t *testing.T) {
	shouldBeInvalid(1111111, t)
}

func shouldBeInvalid(n int64, t *testing.T) {
	result, err := isInvalid(n)
	if err != nil {
		t.Errorf("unexpected error checking if number is invalid: %v", err)
		return
	}
	if !result {
		t.Errorf("%d should be invalid", n)
	}
}
