package main

import "fmt"

func task1() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 != 0 {
			fmt.Println(i)
		} else {
			fmt.Println("Fizz")
		}
	}
}
