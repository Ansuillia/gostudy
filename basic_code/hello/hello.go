package main

import (
	"ansuillia/gostudy/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Ansuillia")
	fmt.Println(message)
}
