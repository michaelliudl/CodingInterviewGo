package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, v := range nums {
		strs[i] = strconv.Itoa(v)
	}
	sort.Slice(strs, func(i, j int) bool {
		if v1, e1 := strconv.Atoi(strs[i] + strs[j]); e1 == nil {
			if v2, e2 := strconv.Atoi(strs[j] + strs[i]); e2 == nil {
				return v1 > v2
			}
		}
		return false
	})
	if strs[0] == "0" {
		return "0"
	}
	return strings.Join(strs, "")
}
