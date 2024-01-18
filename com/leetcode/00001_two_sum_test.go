package main

import "testing"

func TestTwoSum(t *testing.T) {
	t.Logf("Test %d", twoSum([]int{2, 7, 11, 15}, 9))
	t.Logf("Test %d", twoSum([]int{3, 2, 4}, 6))
	t.Logf("Test %d", twoSum([]int{3, 3}, 6))
}

func TestTwoSumWithMap(t *testing.T) {
	t.Logf("Test %d", twoSumWithMap([]int{2, 7, 11, 15}, 9))
	t.Logf("Test %d", twoSumWithMap([]int{3, 2, 4}, 6))
	t.Logf("Test %d", twoSumWithMap([]int{3, 3}, 6))
}
