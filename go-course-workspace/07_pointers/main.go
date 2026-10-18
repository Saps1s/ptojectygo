package main

import "fmt"

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func reverse(arr *[]int) {
	a, b := 0, len(*arr)-1
	for a < b {
		(*arr)[a], (*arr)[b] = (*arr)[b], (*arr)[a]
		a++
		b--
	}
}

type Person struct {
	Name string
	Age  int
}

func celebrateBirthday(p *Person) {
	p.Age++
}

func main() {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)

	vasya := Person{Name: "Василий", Age: 30}
	celebrateBirthday(&vasya)
	fmt.Printf("Возраст: %d\n", vasya.Age)

	arr := []int{10, 20, 30, 40, 50}
	reverse(&arr)
	fmt.Println("После reverse:", arr)
}