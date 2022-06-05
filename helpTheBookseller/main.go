package main

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	var result []string

	countTotal := make(map[string]int)

	for _, art := range listArt {
		splited := strings.Split(art, " ")

		codeLetter := string(splited[0][0])
		codeCount, _ := strconv.Atoi(splited[1])

		if _, exists := countTotal[codeLetter]; exists {
			countTotal[codeLetter] += codeCount
		} else {
			countTotal[codeLetter] = codeCount
		}

	}

	for _, code := range listCat {
		total := 0

		if _, exists := countTotal[code]; exists {
			total = countTotal[code]
		}

		result = append(result, fmt.Sprintf("(%v : %d)", code, total))

	}

	return strings.Join(result, " - ")
}

func main() {

	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}
	fmt.Println(StockList(b, c), " === ", "(A : 0) - (B : 1290) - (C : 515) - (D : 600)")
}

// b = []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"}
// c = []string{"A", "B"}
// dotest(b, c, "(A : 200) - (B : 1140)")
