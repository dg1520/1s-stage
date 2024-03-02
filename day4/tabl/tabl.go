package tabl

import "fmt"

func Tabl() {
	for i := 1; i <= 10; i++{
		for b := 1; b <= 10; b++{
			fmt.Println(i,b, i*b)

		}
		fmt.Println()
	}
	
}


