package foursum

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	sort.Ints(nums)

	ret := make([][]int, 0)
	var last int

	for i := 0; i < len(nums)-3; i++ {
		if nums[i] > target && nums[i] >= 0 {
			break
		}

		if i > 0 && nums[i] == last {
			continue
		} else {
			last = nums[i]
		}

		trips := threeSum(nums[i+1:], target-nums[i])

		for _, trip := range trips {
			ret = append(ret, append([]int{nums[i]}, trip...))
		}
	}

	return ret
}

func threeSum(nums []int, target int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	ret := make([][]int, 0)
	var last int

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > target && nums[i] >= 0 {
			break
		}

		if i > 0 && nums[i] == last {
			continue
		} else {
			last = nums[i]
		}

		pairs := twoSum(nums[i+1:], target-nums[i])

		for _, pair := range pairs {
			ret = append(ret, append([]int{nums[i]}, pair...))
		}
	}

	return ret
}

func twoSum(nums []int, target int) [][]int {
	if len(nums) < 2 {
		return [][]int{}
	}

	ret := make([][]int, 0)
	last := nums[0]

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > target && nums[i] >= 0 {
			break
		}

		if i > 0 && nums[i] == last {
			continue
		} else {
			last = nums[i]
		}

		var lastRight int

		for j := len(nums) - 1; j > i; j-- {
			if j < len(nums)-1 && nums[j] == lastRight {
				continue
			} else {
				lastRight = nums[j]
			}

			if nums[i]+nums[j] == target {
				ret = append(ret, []int{nums[i], nums[j]})
				continue
			}
		}

	}

	return ret
}
