package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	x := sum(1, 2)

	if x != 3 {
		t.Errorf("The sum must be 3, received: %d", x)
	}
}
