package main

import "fmt"

func main() {
	fmt.Println(
		func(a, b int) int {
			return a + b
		}(2, 3))
}

/*
func main() {
	fmt.Println(add(2,3))
}
func add(a, b int) int {
	return a + b
}
*/
