package main

import (
	"runtime/debug"
	"testing"
)

func test(t *testing.T, ptr *int, testFn func(byte) bool, alg func() string) {

	*length = 50
	*ptr = int(r.Next() % uint64(*length))

	pw := alg()
	n := 0
	for i := range pw {
		debug.PrintStack()
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
		test(t, special, IsSpecial, a)
	}
}

func TestDigits(t *testing.T) {
	for _, a := range knownAlgorithms {
		test(t, digits, IsSpecial, a)
	}
}

func TestUppers(t *testing.T) {
	for _, a := range knownAlgorithms {
		test(t, upper, IsSpecial, a)
	}
}

func TestLower(t *testing.T) {
	for _, a := range knownAlgorithms {
		test(t, lower, IsSpecial, a)
	}
}
