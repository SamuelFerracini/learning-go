package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
    message := greetings.Hello("Samuel")
    fmt.Println(message)
}