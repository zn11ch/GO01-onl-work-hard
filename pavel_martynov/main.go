package main

import "fmt"

/*
Создадим простую программу, в которой будет
4 функций, которые по цепочке будут вызывать друг
друга и увеличивать переменную на 1
(просто операция +). Везде использовать переменные,
тоже самое касается и индекса увеличения
(пока мы не смотрели константу, поэтому достаточно
глобальной переменной).
*/

var (
	i int8
)

func fourth() {
	fmt.Println("Fourth function, i ==", i)
	i++
}

func third() {
	fmt.Println("Third function, i ==", i)
	i++

	fourth()
}

func second() {
	fmt.Println("Second function, i ==", i)
	i++

	third()
}

func first() {
	fmt.Println("First function, i ==", i)
	i++

	second()
}

func main() {
	first()

	fmt.Println("After all functions i ==", i)
}
