package main

import "fmt"

func main() {
	fmt.Println(swap())
}

func swap() []int {
	a, b, c := 15, 10, 5
	c, b, a = b, a, c
	return []int{a, b, c}
}
