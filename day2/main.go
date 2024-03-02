package main

import (
	"fmt"
	"log"
)

func main() {
	Age()
	Perimeterp()
	Perimeter()
	Ploshad()

}

func Age() {
	log.Println("Программа запущена")

	var age int
	fmt.Print("Введите ваш возраст: ")
	fmt.Scanln(&age)
	vozrast := 10 - age%10

	fmt.Println("До вашего следующего юбилея вам нужно ещё", vozrast, "лет")

}

func Perimeterp() {

	var length, width float64
	fmt.Print("Введите длину прямоугольника: ")
	fmt.Scanln(&length)
	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Scanln(&width)
	perimeter := 2 * (length + width)
	fmt.Println("Периметр прямоугольника равен", perimeter)
}

func Perimeter() {

	var side float64
	fmt.Print("Введите длину стороны квадрата: ")
	fmt.Scanln(&side)
	perimeter := 4 * side
	fmt.Println("Периметр квадрата равен", perimeter)
}

func Ploshad() {

	var length, width float64
	fmt.Print("Введите длину прямоугольника: ")
	fmt.Scanln(&length)
	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Scanln(&width)
	area := length * width
	fmt.Println("Площадь прямоугольника равна", area)

	log.Println("Программа завершена")
}
