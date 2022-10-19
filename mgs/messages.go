package mgs

import "fmt"

// CheckErr ...
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// PrintMessage ...
func PrintMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}
