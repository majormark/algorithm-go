package stack_and_queue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMaxTree(arr []int) *TreeNode {
	nArr := make([]*TreeNode, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		nArr = append(nArr, &TreeNode{
			Val: arr[i],
		})
	}
	lBigMap := map[*TreeNode]*TreeNode{}
	rBigMap := map[*TreeNode]*TreeNode{}
	stack := make([]*TreeNode, 0, len(arr))
	for _, node := range nArr {
		for len(stack) > 0 && stack[len(stack)-1].Val < node.Val {
			popStackAndSetMap(stack, lBigMap)
		}
		stack = append(stack, node)
	}
	for len(stack) > 0 {
		popStackAndSetMap(stack, lBigMap)
	}
	for _, node := range nArr {
		for len(stack) > 0 && stack[len(stack)-1].Val < node.Val {
			popStackAndSetMap(stack, rBigMap)
		}
		stack = append(stack, node)
	}
	for len(stack) > 0 {
		popStackAndSetMap(stack, rBigMap)
	}
	var head *TreeNode
	for i := 0; i < len(nArr); i++ {
		node := nArr[i]
		left := lBigMap[node]
		right := rBigMap[node]
		if left == nil && right == nil {
			head = node
		} else if left == nil {
			if right.Left == nil {
				right.Left = node
			} else {
				right.Right = node
			}
		} else if right == nil {
			if left.Left == nil {
				left.Left = node
			} else {
				left.Right = node
			}
		} else {
			if left.Val < right.Val {
				if left.Left == nil {
					left.Left = node
				} else {
					left.Right = node
				}
			} else {
				if right.Left == nil {
					right.Left = node
				} else {
					right.Right = node
				}
			}
		}
	}
	return head

}

func popStackAndSetMap(stack []*TreeNode, m map[*TreeNode]*TreeNode) {
	popNode := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	if len(stack) == 0 {
		m[popNode] = nil
	} else {
		m[popNode] = popNode
	}
}
