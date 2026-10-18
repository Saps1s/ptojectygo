package main

import "fmt"

func calculateStats(numbers []int) (int, float64) {
	if len(numbers) == 0 {
		return 0, 0.0
	}
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum, float64(sum) / float64(len(numbers))
}

func findElement(slice []string, value string) (int, bool) {
	for i, v := range slice {
		if v == value {
			return i, true
		}
	}
	return -1, false
}

func factorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func fibonacci(n int) int {
	a, b, i := 0, 1, 1
	for i < n {
		b = a + b
		a = b - a
		i++
	}
	return a
}

func main() {
	fmt.Println("--- Лабораторная ---")
	sum, avg := calculateStats([]int{10, 7, 2})
	fmt.Printf("Сумма: %v, Среднее: %v\n", sum, avg)

	idx, found := findElement([]string{"Marlen", "Maria", "Ilona"}, "Maria")
	fmt.Printf("Индекс: %d, Найден: %t\n", idx, found)

	fmt.Println("--- Факториал ---")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}

	fmt.Println("--- Фибоначчи ---")
	for i := 1; i <= 10; i++ {
		fmt.Printf("№ %d = %d\n", i, fibonacci(i))
	}
}