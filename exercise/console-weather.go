package main

import "fmt"

func main() {
	fmt.Println(weatherCelsius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelsius(34, "Tak sunny"))
	fmt.Println(weatherCelsius(17, "Phuket rainy"))
	fmt.Println(weatherCelsius(9, "Chiang-mai cold"))
}

func weatherCelsius(celsius int, msg string) string {

	tenth := celsius/10
	oneth := celsius%10


	
	display := ""

	if celsius == 25 {
		display = display + " _  _\n"
		display = display + " _||_\n"
		display = display + "|_  _|"
	}

	if celsius == 34 {
		display = display + " _\n"
		display = display + " _||_|\n"
		display = display + " _|  |"
	}

	if celsius == 17 {
		display = display + "    _\n"
		display = display + "  |  |\n"
		display = display + "  |  |"
	}

	if celsius == 9 {
		display = display + " _\n"
		display = display + "|_|\n"
		display = display + " _|"
	}

	display = display + " c\n"
	display = display + msg

	return display
}