package main

import "fmt"

func task2() {
	var even int = 0
	var odd int = 0
	var i int
	for i != -1 {
		fmt.Scanln(&i)
		if i == -1 {
			break
		}
		if i%2 != 0 {
			odd = odd + i
		} else {
			even = even + i
		}

	}
	fmt.Println("even numbers sum is", even)
	fmt.Println("odd numbers sum is", odd)
}
