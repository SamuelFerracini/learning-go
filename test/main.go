package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    names := []string{"Sa", "Sam", "Samu"}

    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }


    for person, message := range messages {
        fmt.Printf("Person: [%s] - Message: [%s]\n", person, message)
    }
}
