package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := []int{}
	for i := 0; i < 10; i++ {
		slice = append(slice, rand.Intn(100)+1)
	}
	fmt.Printf("Исходный массив: %v\n", slice)

	min, max := slice[0], slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Printf("min = %d, max = %d\n", min, max)

	even := []int{}
	for _, v := range slice {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	fmt.Printf("Четные: %v\n", even)
}