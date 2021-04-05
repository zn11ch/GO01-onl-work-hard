package main

import "fmt"

func inc1(i *int) {
	*i = *i + 1

}
func inc2(i *int) {
	*i = *i + 1
}
func inc3(i *int) {
	*i = *i + 1
}
func inc4(i *int) {
	*i = *i + 1

}
func main() {
	var value int = 1
	inc1(&value)
	fmt.Println(value)

	inc2(&value)
	fmt.Println(value)
	inc3(&value)

	fmt.Println(value)
	inc4(&value)
	fmt.Println(value)

}
