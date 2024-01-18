package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var r, c int
	cache := make(map[string]int)
	slow := 0
	for fast := 0; fast < len(s); fast++ {
		fastS := s[fast : fast+1]
		if v, found := cache[fastS]; found {
			cache[fastS] = v + 1
		} else {
			cache[fastS] = 1
		}
		for cache[fastS] > 1 {
			slowS := s[slow : slow+1]
			cache[slowS] = cache[slowS] - 1
			if cache[slowS] == 0 {
				delete(cache, slowS)
			}
			slow++
		}
		c = fast - slow + 1
		if c > r {
			r = c
		}
	}
	return r
}
