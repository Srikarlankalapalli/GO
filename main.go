package main

import (
	"Go/simplecalc"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")

	a, b := 6, 4
	fmt.Println(simplecalc.Add(a, b))

	c, d := 15, 4
	fmt.Println(simplecalc.Sub(c, d))

	e, f := 14.0, 8.8
	fmt.Println(simplecalc.Div(e, f))

	g, h := 24, 8
	fmt.Println(simplecalc.Mult(g, h))

	num := 40
	if num > 40 {
		fmt.Println("Number is greater than 40")
	} else if num == 40 {
		fmt.Println("Number is exactly 40")
	} else {
		fmt.Println("Number is less than 40")
	}

	day := "Wednesday"
	switch day {
	case "Wednesday":
		fmt.Println("Weekday")
	case "Saturday":
		fmt.Println("Weekend")
	default:
		fmt.Println("It's another day")
	}

	fmt.Println("Counting from 1 to 9:")
	for i := 1; i <= 9; i++ {
		fmt.Println(i)
	}

	j := 8
	fmt.Println("Before:", j)
	increase(&j)
	fmt.Println("After:", j)

	defer fmt.Println("Second")
	defer fmt.Println("Middle")
	defer fmt.Println("First")

}

func increase(num *int) {
	*num = *num + 1
}
