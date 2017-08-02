package main

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}

/*
* 我自己的垃圾方法，直接暴力求解
* 时间效率并不高
 */
func twoSum(nums []int, target int) []int {
	var (
		i, vi, k int
		answer   []int
	)
	for i, vi = range nums {
		for k = i + 1; k < len(nums); k++ {
			if vi+nums[k] == target {
				answer = []int{i, k}
			}
		}
	}
	return answer
}

/*
* 大神的方法，采用二分查找法，抓住题干提示（只有唯一解）
 */
func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	f := 0
	half := false
	for i, num := range nums {
		if num*2 == target {
			if half {
				return []int{f, i}
			} else {
				f = i
				half = true
			}
			continue
		}

		m[num] = i
		if v, ok := m[target-num]; ok {
			return []int{v, i}
		}
	}
	return []int{}
}
