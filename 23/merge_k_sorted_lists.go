package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode = nil
	var merged *ListNode = nil
	for {
		min := -1
		for n := range lists {
			if lists[n] != nil {
				if min == -1 {
					min = n
				} else {
					if lists[n].Val < lists[min].Val {
						min = n
					}
				}
			}
		}
		if min == -1 {
			break
		}
		if merged == nil {
			merged = &ListNode{lists[min].Val, nil}
			head = merged
		} else {
			merged.Next = &ListNode{lists[min].Val, nil}
			merged = merged.Next
		}
		lists[min] = lists[min].Next
	}
	return head
}

func mergeKLists2(lists []*ListNode) *ListNode {
	var head *ListNode = &ListNode{}
	var merged *ListNode = head
	for {
		min := -1
		for i, n := range lists {
			if n != nil && min == -1 {
				min = i
			} else if n != nil && n.Val < lists[min].Val {
				min = i
			}
		}
		if min == -1 {
			break
		}
		merged.Next = &ListNode{}
		merged = merged.Next
		merged.Val = lists[min].Val
		lists[min] = lists[min].Next
	}
	return head.Next
}
func mergeKLists3(lists []*ListNode) *ListNode {

	var head *ListNode = &ListNode{}
	var merged *ListNode = head
	var prev *ListNode = nil
	for {
		min := -1
		for i, n := range lists {
			if n != nil && min == -1 {
				min = i
			} else if n != nil && n.Val < lists[min].Val {
				min = i
			}
		}
		if min == -1 {
			if prev != nil {
				prev.Next = nil
			}
			break
		}
		prev = merged
		merged.Val = lists[min].Val
		lists[min] = lists[min].Next
		merged.Next = &ListNode{}
		merged = merged.Next
	}
	if prev != nil {
		return head
	} else {
		return nil
	}
}
