package main

import "math"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	charIndex := make(map[uint8]int)
	var end, start int = 0, 0
	for ; end < n; end++ {
		alpha := s[end]
		if _, ok := charIndex[alpha]; ok {
			start = int(math.Max(float64(charIndex[alpha]), float64(start)))
		}
		ans = int(math.Max(float64(ans), float64(end-start+1)))
		charIndex[s[end]] = end + 1
	}
	return ans
}
