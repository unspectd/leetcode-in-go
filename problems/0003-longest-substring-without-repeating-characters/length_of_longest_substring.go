package _003_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	max, m := 0, map[byte]int{}
	for left, right := 0, 0; right < len(s); right++ {
		if j, ok := m[s[right]]; ok && j >= left {
			left = j + 1
		}
		m[s[right]] = right
		if l := right - left + 1; l > max {
			max = l
		}
	}
	return max
}

func lengthOfLongestSubstringOptimized(s string) int {
	max, m := 0, [128]int{} // 0 = not seen yet, n = seen at index: n-1
	for left, right := 0, 0; right < len(s); right++ {
		if j := m[s[right]]; j > left {
			left = j
		}
		m[s[right]] = right + 1 // setting right + 1 because we store index + 1
		if l := right - left + 1; l > max {
			max = l
		}
	}
	return max
}
