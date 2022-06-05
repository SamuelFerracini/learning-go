package main

import (
	"fmt"
)

func DNAStrand(dna string) string {
	result := ""

	complement := map[string]string{
		"T": "A",
		"A": "T",
		"G": "C",
		"C": "G",
	}

	for _, letter := range dna {
		result += complement[string(letter)]
	}

	return result
}

func main() {
	fmt.Println(DNAStrand("ATTGC"))
}
