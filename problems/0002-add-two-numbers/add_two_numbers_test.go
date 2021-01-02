package _002_add_two_numbers

import (
	"testing"

	. "github.com/unspectd/leetcode-in-go/structs/linked-list"
	"github.com/unspectd/leetcode-in-go/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		gotList1 []int
		gotList2 []int
		want     []int
	}{
		{
			gotList1: []int{2, 4, 3},
			gotList2: []int{5, 6, 4},
			want:     []int{7, 0, 8},
		},
		{
			gotList1: []int{0},
			gotList2: []int{0},
			want:     []int{0},
		},
		{

			gotList1: []int{9, 9, 9, 9, 9, 9, 9},
			gotList2: []int{9, 9, 9, 9},
			want:     []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, testCase := range testCases {
		actual := addTwoNumbers(FromSlice(testCase.gotList1), FromSlice(testCase.gotList2)).ToSlice()

		utils.Checkf(t, is.DeepEqual(actual, testCase.want), testCase)
	}
}
