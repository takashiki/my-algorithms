package tests

import (
	"fmt"
	"testing"

	"github.com/takashiki/my-algorithms/backtrace"
)

func TestPermutation(t *testing.T) {
	nums := []int{1, 2, 3}

	permutations := backtrace.Permute(nums)

	fmt.Printf("%#v\n", permutations)
}

func TestNQueen(t *testing.T) {
	n := 4

	results := backtrace.NQueens(n)

	fmt.Printf("%#v\n", results)
}
