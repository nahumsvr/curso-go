package main

import (
	"fmt"
	"time"
)

func computeFactorial(n int, ch chan string) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- fmt.Sprintf("Factorial de %d es %d", n, result)
}

func main() {
	numbers := []int{2, 3, 4, 5, 6, 7}
	ch := make(chan string, len(numbers))
	for _, n := range numbers {
		go computeFactorial(n, ch)
	}

	for i := 0; i < len(numbers); i++ {
		result := <-ch
		fmt.Println(result)
	}

	close(ch)

	time.Sleep(1 * time.Second) // Wait for goroutines to finish
}
