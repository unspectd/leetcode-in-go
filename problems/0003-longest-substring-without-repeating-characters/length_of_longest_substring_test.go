package _003_longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/unspectd/leetcode-in-go/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "abcabcbb",
			want: 3,
		},
		{
			got:  "bbbbb",
			want: 1,
		},
		{
			got:  "pwwkew",
			want: 3,
		},
	}

	algorithms := []func(string) int{
		lengthOfLongestSubstring,
		lengthOfLongestSubstringOptimized,
	}

	for _, testCase := range testCases {
		for _, algorithm := range algorithms {
			actual := algorithm(testCase.got)

			utils.Checkf(t, is.Equal(actual, testCase.want), testCase)
		}
	}
}
