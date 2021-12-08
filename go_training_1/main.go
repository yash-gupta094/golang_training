package main

import (
	"fmt"
	"strconv"
)

const (
	c, d, i float64 = 50, 3, 23.5
	x               = int32(c * i)
)

var (
	c2, d2 float64 = 6, c2 + i
)

// plain function without parameters
// func main() {
// 	a, b := 10, 30
// 	fmt.Println("Original values: a = ", a, "b = ", b)
// 	// var c int
// 	// c := a
// 	a, b = b, a
// 	fmt.Println("Swapped values: a = ", a, "b = ", b)
// }
func main() {
	// Complex 64
	var r1, i1 float32 = 5.5, 10.5
	var r2, i2 float32 = 6.5, 11.5
	c1 := complex(r1, i1)
	c2 := complex(r2, i2)
	c3 := c1 + c2
	fmt.Println("Complex 64 operations:")
	fmt.Println("C1 + C2 = ", c3)
	fmt.Println("C1  C2 = ", c1-c2)

	fmt.Println("C1 * C2 = ", c1*c2)
	fmt.Println("C1 / C2 = ", c1/c2)

	//Complex 128
	var r641, i641 float64 = 5.5, 10.5
	var r642, i642 float64 = 6.5, 11.5
	c641 := complex(r641, i641)
	c642 := complex(r642, i642)
	c643 := c641 + c642
	fmt.Println("Complex 128 operations:")
	fmt.Println("C1 + C2 = ", c643)
	fmt.Println("C1  C2 = ", c641-c642)

	fmt.Println("C1 * C2 = ", c641*c642)
	fmt.Println("C1 / C2 = ", c641/c642)

	var g float64 = 4.5
	var g2 float32 = 3.3
	fmt.Println(g + float64(g2))
	fmt.Println(strconv.Itoa(int(g)))
	// const d = i + g  Invalid scenario

	fmt.Println("Variable and constant addition:", c2, d2)

	var s int = 5
	var t int = 6
	fmt.Println("Original values: a = ", s, "b = ", t)
	s, t = swap(s, t)
	fmt.Println("Swapped values: a = ", s, "b = ", t)

	// s, _ = swap(s, t)

	// Anonymous functions
	func(str string) {
		println(str)
	}("Hi")

	s, t = func(x int, y int) (int, int) {
		return y, x
	}(s, t)
	fmt.Println("Swapped again: a = ", s, "b = ", t)

	add := calc(s, t, func(x, y int) int {
		return x + y
	})
	fmt.Println("Addition:= ", add)
}

func swap(x int, y int) (int, int) {
	return y, x
}

func calc(a, b int, f func(int, int) int) int {
	return f(a, b)
}
