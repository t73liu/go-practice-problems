package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func factorial(n int32) int32 {
	if n <= 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	result := factorial(n)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}
