package main

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	r := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		n := intervals[i]
		prev := r[len(r)-1]
		if prev[1] < n[0] {
			r = append(r, n)
		} else if prev[1] < n[1] {
			prev[1] = n[1]
		}
	}
	return r
}
