package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	actual := convertToIntArray(strings.Split(readLine(reader), " "))
	expected := convertToIntArray(strings.Split(readLine(reader), " "))

	if expected[2] == actual[2] {
		if expected[1] >= actual[1] {
			if expected[0] >= actual[0] {
				fmt.Println(0)
			} else {
				fmt.Println((actual[0] - expected[0]) * 15)
			}
		} else {
			fmt.Println((actual[1] - expected[1]) * 500)
		}
	} else if expected[2] < actual[2] {
		fmt.Println(10000)
	} else {
		fmt.Println(0)
	}
}
