package greed

import (
	"sort"
	"strconv"
)

func LargestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		// 我一开始的思路是用双指针，循环比较每一位的大小
		// 需要把两个数字转成字符串，比较的时候再转回去，对于强类型语言有点复杂
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		//把两个数拼接起来，比较拼接之后的大小
		return sy*x+y > sx*y+x
	})

	// 排完序后如果第一个数字是 0，说明传入的数组只有一个数字 0
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}

	return string(ans)
}
