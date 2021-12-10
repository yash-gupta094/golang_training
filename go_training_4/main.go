package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	emptyInterfaceExample()

	goRoutineExercise()

	channelExample()

	channelExercise()

	channel2Exercise()

	deferExercise()
}

func deferExercise() {
	s := "hello"
	for _, v := range s {
		defer fmt.Print(string(v))
	}
}
func init() {
	if channel == nil {
		channel = make(chan int)
	}
	fmt.Println("Hello is it called")
}

func channel2Exercise() {

	go Incr2("1", 1)
	go Incr2("2", 1)
	channel <- 1
	time.Sleep(time.Second * 1)
}

var channel chan int

func Incr2(r string, ns int) {
	for {
		fmt.Println("Stil here")
		k := <-channel
		fmt.Println(k)
		if k < 100 {
			k = k + 1
			channel <- k
		}
	}
}

func channelExercise() {
	i := 0
	fmt.Println(i)
	go func(i *int) {
		for {
			if *i >= 100 {
				fmt.Println(*i)
				os.Exit(0)
			}
			fmt.Println("In 1: ", *i)
			*i = *i + 1
		}
	}(&i)
	go func(i *int) {
		for {
			if *i >= 100 {
				fmt.Println(*i)
				os.Exit(0)
			}
			fmt.Println("In 2: ", *i)
			*i = *i + 1
		}
	}(&i)
	time.Sleep(time.Second * 1)
	// go increment(i, ch)
	// go increment(i, ch)
	// fmt.Println(i)
	// if <-ch == true {
	// 	fmt.Println(i)
	// 	os.Exit(0)
	// }
}

func increment(i *int, ch chan bool) int {
	if *i == 100 {
		ch <- true
	}
	return *i + 1
}

func channelExample() {
	var ch chan int
	if ch == nil {
		ch = make(chan int)
	}
	go func() {
		ch <- 100
	}() // Without making the function go routine, there will be a deadlock as there is only one go routine i.e. main
	k := <-ch
	fmt.Println(k)
}

func goRoutineExercise() {
	go hello()
	time.Sleep(time.Millisecond * 1)
	go func() {
		x := 10
		fmt.Println("X = ", x)
	}()
	time.Sleep(time.Millisecond * 1)
	fmt.Println("Caller")
}

func hello() {
	fmt.Println("Hi")
}

func emptyInterfaceExample() {
	fmt.Printf("%f %T\n", Add(1, 2.23, 2), Add(1, 2.23, 2))
	fmt.Printf("%d %T\n", Add(1, 2), Add(1, 2))
}

func Add(values ...interface{}) (sum interface{}) {
	sum = 0
	for _, v := range values {
		switch v.(type) {
		case int:
			switch sum.(type) {
			case float64:
				sum = sum.(float64) + float64(v.(int))
			case int:
				sum = sum.(int) + v.(int)
			}
		case float64:
			switch sum.(type) {
			case int:
				sum = float64(sum.(int))
			}
			sum = sum.(float64) + v.(float64)
		default:
			fmt.Print("Other type")
		}
	}
	return sum
}
