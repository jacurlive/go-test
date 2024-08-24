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
	minNumber := findMin(12, 15, 17, 11, 34, 54, 9)
	fmt.Println(minNumber)
	result := twoSum([]int{3, 2, 4}, 14)
	fmt.Println(result)
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

// for loop, unlimited number of arguments
func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if min > i {
			min = i
		}
	}
	return min
}

// first leetcode problem https://leetcode.com/problems/two-sum/submissions/1366939120/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nums
}
