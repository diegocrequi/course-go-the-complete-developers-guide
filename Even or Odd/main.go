package main

import "fmt"

func main() {
	numbers := []int{}
	for i := 0; i < 11; i++ {
		numbers = append(numbers, i)
	}

	for _, n := range numbers {
		switch n % 2 {
		case 0:
			fmt.Println(n, "is even")
		case 1:
			fmt.Println(n, "is odd")
		}
	}
}
