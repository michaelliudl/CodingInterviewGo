package main

func twoSum(nums []int, target int) []int {
	if nums == nil {
		return []int{}
	}
	var result []int
	for idx := 0; idx < len(nums)-1; idx++ {
		iElem := nums[idx]
		for jdx := idx + 1; jdx < len(nums); jdx++ {
			if nums[jdx] == (target - iElem) {
				return append(append(result, idx), jdx)
			}
		}
	}
	return []int{}
}

func twoSumWithMap(nums []int, target int) []int {
	if nums == nil {
		return []int{}
	}
	var cacheMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cacheMap[nums[i]] = i
	}
	var result []int
	for i := 0; i < len(nums); i++ {
		elem := nums[i]
		if v, ok := cacheMap[target-elem]; ok {
			if v != i {
				return append(append(result, i), v)
			}
		}
	}
	return []int{}
}
