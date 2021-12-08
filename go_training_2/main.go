package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b, c := 8, 1, 2
	findGreatest(a, b, c)

	loopExercise()

	switchExercise()

	fallThroughExercise()

	arrayExercise()

	stringArrayExercise()

	reverseString()

	multiDimensionalExercise()

	sliceExercise()

	variadicExercise()

	mapExercise()

	sliceAndMapExercise()
}

func sliceAndMapExercise() {
	slice := make([]int, 0)
	countMap := map[int]int{}
	for i := 0; i < 100; i++ {
		slice = append(slice, rand.Intn(100))
	}
	fmt.Println(slice)
	for _, value := range slice {
		countMap[value]++
	}
	fmt.Println(countMap)
}

func mapExercise() {
	student := map[int]string{
		1: "Yash",
		2: "Shantanu",
		3: "Anand",
	}
	fmt.Println("Original list of students: ", student)
	addStudent(student, 11, "Mayur")
	fmt.Println("List after student added: ", student)
	updateStudent(student, 11, "Mayuri")
	fmt.Println("Updated list of students : ", student)
	deleteStudent(student, 11)
	fmt.Println("Updated list of students after deletion : ", student)
	fmt.Println("Name of student with roll number:", 2, " is ", getStudent(student, 2))
}

func getStudent(student map[int]string, i int) string {
	return student[i]
}

func deleteStudent(student map[int]string, i int) {
	delete(student, i)
}

func updateStudent(student map[int]string, i int, s string) {
	student[i] = s
}

func addStudent(student map[int]string, i int, s string) {
	if _, ok := student[i]; !ok {
		student[i] = s
		return
	}
	fmt.Println("Can't add student already present with roll no. ", i)
}

func variadicExercise() {
	fmt.Println("Variadic result", add(3, 2, 1, 4, 5))
}

func add(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

//Find greatest in slice
func sliceExercise() {
	slice := make([]int, 0)
	for i := 0; i < 10; i++ {
		slice = append(slice, rand.Intn(900))
	}
	var greatest int = slice[0]
	for _, value := range slice {
		if value > greatest {
			greatest = value
		}
	}
	fmt.Println("Slice : ", slice)
	fmt.Println("Greatest value in slice is: ", greatest)

	// Remove element from index 3 of slice
	slice = append(slice[:3], slice[3+1:]...)
	fmt.Println("Removed third element from slice: ", slice)
}

func multiDimensionalExercise() {
	arr := [2][2][2]int{
		{
			{1, 2},
			{2, 3},
		},
		{
			{4, 5},
			{5, 6},
		},
	}
	arr2 := [2][2][2]int{
		{
			{11, 12},
			{12, 13},
		},
		{
			{14, 15},
			{15, 16},
		},
	}
	var result [2][2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				result[i][j][k] = arr[i][j][k] + arr2[i][j][k]
			}
		}

	}
	fmt.Println("Addition of matrices: ", result)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				result[i][j][k] = arr2[i][j][k] - arr[i][j][k]
			}
		}
	}
	fmt.Println("Subtraction of matrices: ", result)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				result[i][j][k] = arr2[i][j][k] * arr[i][j][k]
			}
		}
	}
	fmt.Println("Multiplication of matrices: ", result)

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
