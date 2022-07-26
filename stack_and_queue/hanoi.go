package stack_and_queue

import "fmt"

func hanoiV1(num int, A string, B string, C string) {
	if num == 1 {
		move(num, A, C)
	}
	hanoiV1(num-1, A, C, B)
	move(num, A, B)
	hanoiV1(num-1, B, A, C)
}

func move(num int, from string, to string) {
	fmt.Println(fmt.Sprintf("Move %v from %v to %v.", num, from, to))
}

func hanoiV2(num int, left string, mid string, right string, from string, to string) {
	if num == 1 {
		if from == mid || to == mid {
			move(1, from, to)
		} else {
			move(1, from, mid)
			move(1, mid, to)
		}
	}
	if from == mid || to == mid {
		another := left
		if from == left || to == left {
			another = right
		}
		hanoiV2(num-1, left, mid, right, from, another)
		move(num, from, to)
		hanoiV2(num-1, left, mid, right, another, to)
	} else {
		hanoiV2(num-1, left, mid, right, from, to)
		move(num, from, mid)
		hanoiV2(num-1, left, mid, from, to, from)
		move(num, mid, to)
		hanoiV2(num-1, left, mid, right, from, to)
	}

}
