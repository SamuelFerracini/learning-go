package main

import (
	"fmt"
	"strconv"
)

func DigitalRoot(n int) int {

	str := strconv.Itoa(n)

	if len(str) == 1 {
		return n
	}

	result := 0

	for _, digit := range str {

		xint, _ := strconv.Atoi(string(digit))

		result += xint
	}

	return DigitalRoot(result)

}

func main() {
	fmt.Println(DigitalRoot(13242342310))
}
