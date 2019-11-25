package main

import (
	"fmt"
	"testing"
)

// expected output is 3
func TestMinimumSwaps1(t *testing.T) {
	arr := []int32{4, 3, 1, 2}

	swaps := minimumSwaps(arr)

	fmt.Println("SWAPS:", swaps)
}

// expected output is
func TestMinimumSwaps2(t *testing.T) {
	arr := []int32{1, 3, 5, 2, 4, 6, 7}

	swaps := minimumSwaps(arr)

	fmt.Println("SWAPS:", swaps)
}
