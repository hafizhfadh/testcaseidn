package main

import (
	"fmt"
	"idn/calculation"
)

func main() {
	// Answer Number One //

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// Answer Number Two //

	elems := 1000
	var array = make([]int, elems)
	for i := 1; i < 1000; i++ {
		if i%3 == 0 {
			array[i] = i
		} else if i%5 == 0 {
			array[i] = i
		}
	}
	fmt.Println(calculation.Sum(array))
}
