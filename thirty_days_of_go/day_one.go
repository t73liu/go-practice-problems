package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d = 4.0
	var s = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var newInt int64
	var newDouble float64
	var newText string

	// Read and save an integer, double, and String to your variables.
	scanner.Scan()
	newInt, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	newDouble, _ = strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	newText = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(newInt + int64(i))

	// Print the sum of the double variables on a new line.
	fmt.Println(strconv.FormatFloat(d+newDouble, 'f', 1, 64))

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + newText)
}
