package main

import "testing"

func TestMergeInternvals(t *testing.T) {
	// t.Logf("Test %d", mergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	// t.Logf("Test %d", mergeIntervals([][]int{{1, 4}, {4, 5}}))
	t.Logf("Test %d", mergeIntervals([][]int{{1, 4}, {0, 4}}))
	t.Logf("Test %d", mergeIntervals([][]int{{1, 4}, {2, 3}}))
}
