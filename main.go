package main

import (
	"fmt"
	"os"
)

func main() {
	const kumar = 10
	arr := [2]string{"kumar", "our"}
	// kumar = 11
	fmt.Println(kumar, arr)

	f, err := os.Open("filename.ext")
	if err != nil {
		fmt.Println("kumar", f)
	}
	if err != nil {
		fmt.Printf("kumar oru po***.")
	}
	var number int
	fmt.Println("What is your favorite number?")
	fmt.Scan(&number)
	fmt.Printf("Your favorite number is %d.\n", number)
}
