package fib

import "fmt"

func Fib() {
	for i := 0; i <= 10; i++ {
		var d int
		d = (i - 1) + (i - 2)

		fmt.Println(i+i, d, i)
	}

}
