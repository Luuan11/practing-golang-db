package main

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}}

	for i, test := range tests {
		result := fibonacci(test.arg)
		if result != test.want {
			t.Errorf("Test[%d]: fibonacci length(%d) returned %d, want %d", i, test.arg, result, test.want)
		}
	}
}
