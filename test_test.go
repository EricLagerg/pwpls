package main

import "testing"

func test(t *testing.T, ptr *int, testFn func(byte) bool, alg func(bool) string) {

	*length = 50
	*ptr = int(r.Next() % uint64(*length))

	pw := alg(false)
	n := 0
	for i := range pw {
		if testFn(pw[i]) {
			n++
		}
	}

	if n != *ptr {
		t.Errorf("Wanted %d, got %d special characters", *ptr, n)
	}
}

func TestSpecs(t *testing.T) {
	for _, a := range knownAlgorithms {
		test(t, special, isSpecial, a)
	}
}

func TestDigits(t *testing.T) {
	for _, a := range knownAlgorithms {
		test(t, digits, isSpecial, a)
	}
}

func TestUppers(t *testing.T) {
	for _, a := range knownAlgorithms {
		test(t, upper, isSpecial, a)
	}
}

func TestLower(t *testing.T) {
	for _, a := range knownAlgorithms {
		test(t, lower, isSpecial, a)
	}
}
