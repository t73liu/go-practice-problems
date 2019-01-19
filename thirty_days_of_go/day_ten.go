package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	n := int64(nTemp)
	nBinary := strconv.FormatInt(n, 2)

	max := 0
	curr := 0

	for _, char := range nBinary {
		if string(char) == "1" {
			curr += 1
		} else {
			curr = 0
		}
		if curr > max {
			max = curr
		}
	}

	fmt.Println(max)
}
