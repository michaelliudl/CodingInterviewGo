package main

import "testing"

func TestLargestNumber(t *testing.T) {
	t.Logf("Test %s", largestNumber([]int{10, 2}))
	t.Logf("Test %s", largestNumber([]int{3, 30, 34, 5, 9}))
}
