package main

func f(n int) bool {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum == (n*(n+1))/2
}
