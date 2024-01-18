package main

import "testing"

func TestRemoveElement(t *testing.T) {
	t.Logf("Test %d", removeElement([]int{3, 2, 2, 3}, 3))
	t.Logf("Test %d", removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
