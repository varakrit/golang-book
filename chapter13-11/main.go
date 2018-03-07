package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3) // buffer 3
	go func() { ch <- 1 }()
	ch <- 2
	fmt.Println("cap:", cap(ch)) // buffer capacity
	fmt.Println("len:", len(ch)) // buffer used
}
