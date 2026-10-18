package main

import "fmt"

type Book struct {
	ID       int
	Title    string
	Author   string
	IsIssued bool
}

type Reader struct {
	ID   int
	Name string
}

func main() {
	fmt.Println("Проект 'Простая библиотека' запущен.")
	b := Book{ID: 1, Title: "Война и мир", Author: "Лев Толстой"}
	r := Reader{ID: 1, Name: "Иван Иванов"}
	fmt.Println(b, r)
}