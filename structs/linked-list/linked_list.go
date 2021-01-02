package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func FromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	first := &ListNode{Val: nums[0]}
	for current, i := first, 1; i < len(nums); i++ {
		next := ListNode{Val: nums[i]}
		current.Next = &next
		current = &next
	}
	return first
}

func (list *ListNode) ToSlice() []int {
	if list == nil {
		return []int{}
	}

	var slice []int
	for list != nil {
		slice = append(slice, list.Val)
		list = list.Next
	}
	return slice
}
