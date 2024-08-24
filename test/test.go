package main

import (
	"errors"
	"fmt"
	"log"
)

// every programm start here (main function - is first step of every programm)
func main() {
	message, error := enterTheClub(18)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)
	prediction := prediction("пн")
	fmt.Println(prediction)
}

// if else function
func enterTheClub(age int) (string, error) {
	if age >= 18 {
		return "Welcome", nil
	}
	return "STOP!", errors.New("you are too young")
}

// switch case function
func prediction(dayOfWeek string) string {
	switch dayOfWeek {
	case "пн":
		return "Понедельник"
	case "вт":
		return "Вторник"
	default:
		return "No prediction"
	}

}
