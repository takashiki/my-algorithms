package backtrace

var permutations [][]int

func Permute(nums []int) [][]int {
	track := make([]int, 0)
	permuteBacktrace(nums, track)

	return permutations
}

func permuteBacktrace(nums []int, track []int) {
	if len(track) == len(nums) {
		permutations = append(permutations, track)
	}

	for _, x := range nums {
		if contains(track, x) {
			continue
		}

		track = append(track, x)

		permuteBacktrace(nums, track)

		track = track[:len(track)-1]
	}
}

func contains(a []int, n int) bool {
	for _, d := range a {
		if d == n {
			return true
		}
	}

	return false
}
