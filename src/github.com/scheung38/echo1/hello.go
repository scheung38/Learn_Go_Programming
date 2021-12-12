package main

import (
	"fmt"
	"math"
	// "os"
)

// func hypot(x, y float64) float64 {
// 	return math.Sqrt(x*x + y*y)
// }

func add2(x float64, y float64) float64 {
	return math.Sqrt(x + y)
}

func main() {

	res := add2(80, 1)

	s := fmt.Sprintf("%f", res)

	fmt.Println(s)

	// ch := make(chan int)

	// ch <- 5

	// value := <- ch

	fmt.Println("Hello, Sebastian 世界")

	// var s, sep string
	// for i:=1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	// var k = 55
	// var l = 67.8

	// sum := k + int(l)
	// fmt.Printf(string(sum))

}
