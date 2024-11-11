package main

import (
	"fmt"
)

func sum(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func product(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}

func subtraction(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]             // start with the first value
	for _, v := range i[1:] { // subtract the remaining values
		total -= v
	}
	return total
}

func division(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0] // start with the first number
	for _, v := range i[1:] {
		if v == 0 {
			fmt.Println("Error: division by zero")
			return 0
		}
		total /= v

	}
	return total
}

func main() {
	x := sum(1, 2, 3)
	y := product(10, 10)
	z := subtraction(10, 3, 2)
	w := division(4 / 2)

	fmt.Println(" sum result:", x, "\n product result:", y, "\n subtraction result:", z, "\n division result:", w)
}
