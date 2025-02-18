package project

import "fmt"

//go:generate go run gen.go

func PrintContributors() {
	for _, contributor := range Contributors {
		fmt.Println(contributor)
	}
}
