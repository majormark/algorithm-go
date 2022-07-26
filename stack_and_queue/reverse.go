package stack_and_queue

func getAndRemoveLastElement(stack []int) int {
	res := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	if len(stack) == 0 {
		return res
	} else {
		last := getAndRemoveLastElement(stack)
		stack = append(stack, res)
		return last
	}
}

func reverse(stack []int) {
	if len(stack) <= 1 {
		return
	}
	last := getAndRemoveLastElement(stack)
	reverse(stack)
	stack = append(stack, last)
}
