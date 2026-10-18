package main

import "fmt"

func main() {
	studentName := "Иван"
	studentAge := 20
	isStudent := true

	fmt.Printf("Имя: %v, Возраст: %v, Является студентом: %v\n",
		studentName, studentAge, isStudent)

	a := 15
	b := 5
	fmt.Println("Сумма:", a+b)
	fmt.Println("Разность:", a-b)
	fmt.Println("Произведение:", a*b)
	fmt.Println("Частное:", a/b)
}