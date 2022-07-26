package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[:3]
	b = append(b, 1)
	fmt.Println(b)
	fmt.Println(a)
}
