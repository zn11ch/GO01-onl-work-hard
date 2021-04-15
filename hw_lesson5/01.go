package main

import "fmt"

func fibIterative(n int) int {
	a, b := 0, 1
	for i := 1; i <= n; i++ {
		a, b = b, a+b
	}
	return a
}

func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

func fibClosure() func() int {
	a, b := 0, 1
	return func() int {
		res := a
		a, b = b, a+b
		return res
	}
}

func finChannel(n int, c chan int) {
	a, b := 0, 1
	for i := 1; i <= n; i++ {
		c <- a
		a, b = b, a+b
	}
	close(c)
}

func main() {
	n := 10
	fmt.Print("Iterative: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", fibIterative(i))
	}
	fmt.Println()

	fmt.Print("Recursive: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", fibRecursive(i))
	}
	fmt.Println()

	fmt.Print("Closure: ")
	nextFib := fibClosure()
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", nextFib())
	}
	fmt.Println()

	fmt.Print("Channel: ")
	c := make(chan int)
	go finChannel(n, c)
	for val := range c {
		fmt.Printf("%d ", val)
	}
}
