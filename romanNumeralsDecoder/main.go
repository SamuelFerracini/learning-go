package main

import (
	"fmt"
	"strings"
)

func Decode(roman string) int {
	value := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	splited := strings.Split(roman, "")

	result := 0

	for idx, letter := range splited {
		if idx+1 < len(splited) && value[letter] < value[splited[idx+1]] {
			result -= value[letter]
		} else {
			result += value[letter]
		}
	}

	return result
}

func main() {
	fmt.Println(Decode("MDCLXVI"))
}
