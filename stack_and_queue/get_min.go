package stack_and_queue

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) > 0 {
		top := this.minStack[len(this.minStack)-1]
		if top <= val {
			this.minStack = append(this.minStack, top)
		} else {
			this.minStack = append(this.minStack, val)
		}
	} else {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]

}
