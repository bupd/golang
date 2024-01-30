package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int8 = 127
	// intNum += 0.2
	fmt.Println(intNum)

	// kumar := 0.33333333333
	ku := 2 / 3
	ku1 := 1 / 3
	fmt.Println(ku, ku1)

	var kumaruu float64 = float64(intNum) + 0.5

	fmt.Println(ku1+ku, kumaruu)
	Stringoo := "🐜"

	fmt.Println(len(Stringoo)) // this shows the len in BYTES
	fmt.Println(utf8.RuneCountInString(Stringoo))
}
