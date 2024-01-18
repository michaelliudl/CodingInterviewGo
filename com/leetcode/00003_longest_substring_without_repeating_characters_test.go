package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Logf("Test %d", lengthOfLongestSubstring("abcabcbb"))
	t.Logf("Test %d", lengthOfLongestSubstring("bbbbb"))
	t.Logf("Test %d", lengthOfLongestSubstring("pwwkew"))
	t.Logf("Test %d", lengthOfLongestSubstring(" "))
}
