package stack_and_queue

type MyQueue struct {
	stackPush []int
	stackPop  []int
}

func MyQueueConstructor() MyQueue {
	return MyQueue{
		stackPush: []int{},
		stackPop:  []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stackPush = append(this.stackPush, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stackPop) == 0 {
		for i := len(this.stackPush) - 1; i >= 0; i-- {
			this.stackPop = append(this.stackPop, this.stackPush[i])
		}
		this.stackPush = this.stackPush[:0]
	}
	val := this.stackPop[len(this.stackPop)-1]
	this.stackPop = this.stackPop[:len(this.stackPop)-1]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.stackPop) == 0 {
		for i := len(this.stackPush) - 1; i >= 0; i-- {
			this.stackPop = append(this.stackPop, this.stackPush[i])
		}
		this.stackPush = this.stackPush[:0]
	}
	return this.stackPop[len(this.stackPop)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stackPush) == 0 && len(this.stackPop) == 0
}
