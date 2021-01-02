package linked_list

import (
	"testing"

	"github.com/unspectd/leetcode-in-go/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestFromToSlice(t *testing.T) {
	testCases := [][]int{
		{},
		{0, 1, 4, 8},
	}

	for _, testCase := range testCases {
		list := FromSlice(testCase)
		slice := list.ToSlice()

		utils.Checkf(t, is.DeepEqual(slice, testCase), testCase)
	}
}
