package main

import "time"
import "fmt"
import "math/rand"

func main() {

	var target int
	var input int

	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

	target=r1.Intn(10)
	for number:=1; number < 7; number++ {
		fmt.Print("Please guess a number: ")
		fmt.Scanln(&input)
		if number == 6 {
			fmt.Println("เกินพอ")
			fmt.Println("คำตอบคือ",target)
		} else if input > target {
			fmt.Println("มากไป")
		} else if input < target {
			fmt.Println("น้อยไป")
		} else {
			fmt.Println("เจอแล้ว")
			number = 10
		}
	}
}
