package main

import "fmt"

func factIter(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func factRecurs(n int) int {
	if n == 1 {
		return 1
	}
	return n * factRecurs(n-1)
}

func main() {
	n := 5
	fmt.Println(factIter(n))
	fmt.Println(factRecurs(n))
}
