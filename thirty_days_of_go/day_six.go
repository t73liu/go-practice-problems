package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	iterations, _ := strconv.ParseInt(readLine(reader), 10, 64)

	for i := int64(1); i <= iterations; i++ {
		line := readLine(reader)
		firstHalf := ""
		secondHalf := ""
		for pos, char := range line {
			if pos%2 == 0 {
				firstHalf += string(char)
			} else {
				secondHalf += string(char)
			}
		}
		fmt.Println(firstHalf + " " + secondHalf)
	}
}
