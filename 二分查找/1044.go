package binarysearch

func longestDupSubstring(s string) string {
	left, right := 0, len(s)-1
	ans := ""
	for left < right {
		mid := left + (right-left+1)>>1
		if res, ok := getMax(s, mid); ok {
			ans = res
			left = mid
		} else {
			right = mid - 1
		}
	}
	return ans
}

func getMax(s string, length int) (string, bool) {
	cache := map[string]struct{}{}
	for i := -0; i <= len(s)-length; i++ {
		subStr := s[i : i+length]
		if _, ok := cache[subStr]; ok {
			return subStr, true
		}
		cache[subStr] = struct{}{}
	}
	return "", false
}
