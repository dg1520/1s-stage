package main

import (
	"day/day4/fib"
	"day/day4/tabl"
	"fmt"
	"log"
)

func main() {
	Tabl()
	Fib()
	day1()
	day2()
	day3()
	day4()
	log.Println("Программа запущена")

	fmt.Println("Здравствуйте! Пожалуйста, выберите день обучения:")
	fmt.Println("- 1 День. Знакомство")
	fmt.Println("- 2 День. Первая программа")
	fmt.Println("- 3 День. Типы, Функции и Модули")
	fmt.Println("- 4 День. Условия и Циклы")

	var day int
	fmt.Print("Введите число дня (1, 2 или 3): ")
	fmt.Scan(&day)

	switch day {
	case 1:
		day1()
	case 2:
		day2()
	case 3:
		day3()
	case 4:
		day4()
	default:
		fmt.Println("Домашнего задания для выбранного дня нет.")
	}

}

func day1() {
	fmt.Println("Домашнего задания для выбранного дня нет.")
}

func day2() {
	fmt.Println("Д/З Написать программу, которая находит периметр квадрата. Пользователь должен сам вводить числа. Написать программу, которая находит периметр прямоугольника. Пользователь должен сам вводить числа. Программа, которая находит площадь прямоугольника. Пользователь должен сам вводить числа. Написать программу, которая будет вычислять, сколько пользователю ещё нужно лет до его юбилея.")

	var side float64
	fmt.Print("Введите длину стороны квадрата: ")
	fmt.Scanln(&side)
	perimeter := 4 * side
	fmt.Println("Периметр квадрата равен", perimeter)

	var length, width float64
	fmt.Print("Введите длину прямоугольника: ")
	fmt.Scanln(&length)
	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Scanln(&width)
	perimeter = 2 * (length + width)
	fmt.Println("Периметр прямоугольника равен", perimeter)

	fmt.Print("Введите длину прямоугольника: ")
	fmt.Scanln(&length)
	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Scanln(&width)
	area := length * width
	fmt.Println("Площадь прямоугольника равна", area)

	var age int
	fmt.Print("Введите ваш возраст: ")
	fmt.Scanln(&age)
	youbil := 10 - age%10
	fmt.Println("До вашего следующего юбилея вам нужно ещё", youbil, "лет")

}

func day3() {
	fmt.Println("Д/З Необходимо блоки кода разбить на функции. Распределите код из main.go на пакеты и отдельные Go файлы. А в main только вызывайте существующие функции. Создайте модуль.")
}

func day4() {
	fmt.Println("Д/З Создать программу которая объединить все д/з в единую программу. Важные условия: - Если на выбранный пользователем день не было домашнего задания, программа должна об этом сообщить пользователю. Если же день или Д/З определённого дня, которые ввёл пользователь нами ещё не пройдены (или такого Д/З не существует), программа так же об этом должна сообщить пользователю. Бонусное задание: При помощи циклов вывести в терминал таблицу умножения. Используя рекурсию написать функцию, вычисляющую числа Фибоначчи.")

	num := 10
	for i := 0; i < num; i++ {
		fmt.Println()
	}
	log.Println("Программа завершена")
}
