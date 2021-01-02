package _001_two_sum

import (
	"testing"

	"github.com/unspectd/leetcode-in-go/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      []int
	}{
		{
			gotNums:   []int{2, 7, 11, 12},
			gotTarget: 19,
			want:      []int{1, 3},
		},
		{
			gotNums:   []int{3, 3},
			gotTarget: 6,
			want:      []int{0, 1},
		},
	}

	for _, testCase := range testCases {
		actual := twoSum(testCase.gotNums, testCase.gotTarget)

		utils.Checkf(t, is.DeepEqual(actual, testCase.want), testCase)
	}
}
