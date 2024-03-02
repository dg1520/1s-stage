package main

import (
 "fmt"
 "math/rand"
)

// Структура "Ученик"
type Student struct {
 name            string
 grade           int
 lessonsAttended int
 hasNotebook     bool
}

// Функция для выставления оценки ученику
func gradeStudent(student Student) int {
 grade := student.grade

 if student.lessonsAttended >= 15 {
  grade++ // Плюс к оценке за посещение уроков
 }

 if student.hasNotebook {
  grade++ // Плюс к оценке за наличие тетради с темами
 }

 return grade
}

func main() {
 rand.Seed(42) // Устанавливаем seed для генерации случайных чисел

 // Создаем учеников
 students := []Student{
  {"Иван", 4, 16, true},
  {"Мария", 3, 12, false},
  {"Петр", 5, 18, true},
 }

 // Экзамен
 for _, student := range students {
  fmt.Printf("%s выбирает билет...\n", student.name)
  // Предположим, что ученик отвечает на вопросы экзамена

  fmt.Printf("%s получил оценку: %d\n", student.name, gradeStudent(student))
 }

 // Драка между учениками
 fightingStudent := students[rand.Intn(len(students))]
 fmt.Printf("На экзамене началась драка между учениками! %s отправлен в больницу.\n", fightingStudent.name)
}
