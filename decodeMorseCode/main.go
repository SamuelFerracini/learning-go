package main

import (
	"fmt"
	"strings"
)

func DecodeMorse(morseCode string) string {
	morse := map[string]string{
		"-----":     "0",
		".----":     "1",
		"..---":     "2",
		"...--":     "3",
		"....-":     "4",
		".....":     "5",
		"-....":     "6",
		"--...":     "7",
		"---..":     "8",
		"----.":     "9",
		".-":        "a",
		"-...":      "b",
		"-.-.":      "c",
		"-..":       "d",
		".":         "e",
		"..-.":      "f",
		"--.":       "g",
		"....":      "h",
		"..":        "i",
		".---":      "j",
		"-.-":       "k",
		".-..":      "l",
		"--":        "m",
		"-.":        "n",
		"---":       "o",
		".--.":      "p",
		"--.-":      "q",
		".-.":       "r",
		"...":       "s",
		"-":         "t",
		"..-":       "u",
		"...-":      "v",
		".--":       "w",
		"-..-":      "x",
		"-.--":      "y",
		"--..":      "z",
		".-.-.-":    ".",
		"--..--":    ",",
		"..--..":    "?",
		"-.-.--":    "!",
		"-....-":    "-",
		"-..-.":     "/",
		".--.-.":    "@",
		"-.--.":     "(",
		"-.--.-":    ")",
		"...---...": "SOS",
	}

	splited := strings.Split(strings.Trim(morseCode, " "), " ")

	var result []string

	for _, signal := range splited {
		if signal == "" && result[len(result)-1] != " " {
			result = append(result, " ")
		} else {
			x := morse[signal]
			result = append(result, strings.ToUpper(x))
		}
	}

	return strings.Join(result, "")
}

func main() {
	result := DecodeMorse("...---... -.-.--   - .... .   --.- ..- .. -.-. -.-   -... .-. --- .-- -.   ..-. --- -..-   .--- ..- -- .--. ...   --- ...- . .-.   - .... .   .-.. .- --.. -.--   -.. --- --. .-.-.-")
	fmt.Println(result)
}

// Expect(DecodeMorse(".... . -.--   .--- ..- -.. .")).To(Equal("HEY JUDE"))
