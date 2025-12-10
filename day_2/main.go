package main

import (
	"adventofcode/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Read File
	content, err := helpers.ReadFile("day_2/input.txt")
	helpers.Check(err)
	all_ranges := strings.Split(content, ",")

	// Declare variables for later
	var invalids []int
	password := 0

	// For all ranges in the file
	for _, r := range all_ranges {
		// Get lower and upper
		lower_upper_ranges := strings.Split(r, "-")
		lower, err := strconv.Atoi(strings.TrimSpace(lower_upper_ranges[0]))
		helpers.Check(err)
		upper, err := strconv.Atoi(strings.TrimSpace(lower_upper_ranges[1]))
		helpers.Check(err)

		// Loop over range r
		for i := lower; i <= upper; i++ {
			// Check for invalid
			if checkIsRepeated(i) || p2check(i) {
				// if invalid add to array
				invalids = append(invalids, i)
				// fmt.Println(i)
			}

		}

		// add invalid ids to password

	}
	for _, invalid := range invalids {
		password += invalid
	}
	fmt.Println(password)
}

func checkIsRepeated(numToCheck int) bool {
	// Convert number to string, to use string splicing
	string_num := strconv.Itoa(numToCheck)

	// check if even length'd id (has to be, for correct answer) and for no leading 0s
	if len(string_num)%2 == 0 && string_num[0] != 0 {
		// get half way point of string
		half_len := len(string_num) / 2

		// Loop over first half - if not equal to same position in second half, not invalid ID, return false
		for i := range half_len {
			if string_num[i] != string_num[i+half_len] {
				return false
			}
		}
		// fmt.Println(string_num)
		return true
	}
	return false
}

func p2check(numToCheck int) bool {
	string_num := strconv.Itoa(numToCheck)

	isValid := false
	for i := range string_num {

		j := i + 1
		fmt.Println(string_num, i, j)
		foundNum := false
		for !foundNum {

			if i <= j && j < len(string_num) {
				// fmt.Println(string_num, string_num[i], string_num[j])
				if string_num[i] == string_num[j] {
					foundNum = true
					break
				}
			}
			if j == len(string_num)-1 {
				break
			} else {
				// fmt.Println("here")
				j++
			}

		}
		for n := range len(string_num[i:j]) - 1 {
			fmt.Println(string_num, i, j, string_num[i:j], n)
			if i+n <= len(string_num) && j+n <= len(string_num) {
				if string_num[i+n] == string_num[j+n] {
					isValid = true
				}
			}

		}

	}
	if isValid {
		fmt.Println(string_num)
	}
	return isValid
}
