package main

import "fmt"

func main() {
	message, enter := enterTheClub(18)
	fmt.Println(message, enter)
}

func enterTheClub(age int) (string, bool) {
	if age >= 18 {
		return "Welcome", true
	}
	return "STOP!", false
}
