package main

import (
	"project/calc"
	"fmt"
)

func main() {
	num1 := 1184
	num2 := 2

	sum := calc.A(num1, num2)
	diff := calc.S(num1, num2)

	fmt.Println("Sum:", sum)
	fmt.Println("Diff:", diff)
}
