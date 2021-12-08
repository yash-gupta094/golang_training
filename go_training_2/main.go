package main

import "fmt"

func main() {
	a, b, c := 8, 1, 2
	findGreatest(a, b, c)

	loopExercise()

	switchExercise()

	fallThroughExercise()
}

func fallThroughExercise() {
	i := 47
	switch {
	case i < 10:
		fmt.Println(i, " is less than 10")
		fallthrough
	case i < 50:
		fmt.Println(i, " is less than 50")
		fallthrough
	case i < 100:
		fmt.Println(i, " is less than 100")
	}
}

func switchExercise() {
	n := 224
	for i := 0; i < n; i++ {
		switch {
		case i%2 == 0 && i%3 == 0 && i%4 == 0:
			fmt.Println(i, " is divisible by 2, 3 and 4")
		case i%5 == 0 && i%6 == 0 && i%7 == 0:
			fmt.Println(i, " is divisible by 5, 6 and 7")
		case i%8 == 0 && i%9 == 0:
			fmt.Println(i, " is divisible by 9 and 8")
		}
	}
}

func loopExercise() {
	n := 5
	for i := n; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func findGreatest(a, b, c int) {
	if a > b && a > c {
		fmt.Println("A is greatest")
	} else if b > a && b > c {
		fmt.Println("B is greatest")
	} else if c > a && c > b {
		fmt.Println("C is greatest")
	} else {
		fmt.Println("Duplicate values present")
	}
}
