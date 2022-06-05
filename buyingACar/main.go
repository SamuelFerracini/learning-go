package main

import (
	"fmt"
	"math"
)

func NbMonths(startPriceOld int, startPriceNew int, savingperMonth int, percentLossByMonth float64) [2]int {

	countMonths := 0
	savedMoney := 0.0

	floatStartPriceOld := float64(startPriceOld)
	floatStartPriceNew := float64(startPriceNew)
	floatSavingperMonth := float64(savingperMonth)

	for (floatStartPriceOld + savedMoney) < floatStartPriceNew {

		floatStartPriceOld -= floatStartPriceOld * (percentLossByMonth / 100)
		floatStartPriceNew -= floatStartPriceNew * (percentLossByMonth / 100)

		if countMonths%2 == 0 {
			percentLossByMonth += 0.5
		}

		savedMoney += floatSavingperMonth
		countMonths++
	}

	return [2]int{countMonths, int(math.Round(floatStartPriceOld + savedMoney - floatStartPriceNew))}

}

func main() {
	result := NbMonths(7500, 32000, 300, 1.55)
	fmt.Println(result)
}

// dotest(2000, 8000, 1000, 1.5, [2]int{6, 766})
// dotest(12000, 8000, 1000, 1.5 , [2]int{0, 4000})
// dotest(8000, 12000, 500, 1.0, [2]int{8, 597})
// dotest(18000, 32000, 1500, 1.25, [2]int{8, 332})
// dotest(7500, 32000, 300, 1.55, [2]int{25, 122})
