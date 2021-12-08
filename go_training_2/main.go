package main

import "fmt"

func main() {
	a, b, c := 8, 1, 2
	findGreatest(a, b, c)

	loopExercise()

	switchExercise()

	fallThroughExercise()

	arrayExercise()

	stringArrayExercise()

	reverseString()
}

func reverseString() {
	x := "Golang"
	var y string = ""
	for _, value := range x {
		y = string(value) + y
	}
	fmt.Println("Reversed string: ", y)
}

func stringArrayExercise() {
	str := "Hey@##"
	noOfVowels := 0
	noOfSpecialCharacters := 0
	noOfConsonants := 0
	for _, value := range str {
		char := string(value)
		switch {
		case char == "a" || char == "e" || char == "i" || char == "o" || char == "u" || char == "A" || char == "E" || char == "I" || char == "O" || char == "U":
			noOfVowels++
		case ((char >= "a" && char <= "z") || (char >= "A" && char <= "Z")):
			noOfConsonants++
		default:
			noOfSpecialCharacters++
		}
	}
	fmt.Println("Vowels: ", noOfVowels, " Consonants: ", noOfConsonants, " Special Characters:", noOfSpecialCharacters)

}

func arrayExercise() {
	num := [...]int{1, 2, 3, 4, 5}
	sum := 0
	for _, value := range num {
		sum += value
	}
	fmt.Println("Sum of array: ", sum)
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
