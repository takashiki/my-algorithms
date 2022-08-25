package o

import (
	"math"

	"github.com/takashiki/my-algorithms/common"
)

func MaxGcd(nums []int) int {
	min, max := math.MaxInt, 0
	for _, n := range nums {
		min = common.Min(min, n)
		max = common.Max(max, n)
	}

	return gcd(min, max)
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}

	r := a % b
	if r == 0 {
		return b
	}

	return gcd(r, b)
}
