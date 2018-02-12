package main

import "fmt"

func main() {
	fmt.Print("Enter a Feet: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input*0.3048
	fmt.Printf("%.2f", output)
}