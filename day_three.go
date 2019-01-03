package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)

	if N%2 == 1 {
		fmt.Println("Weird")
	} else if 2 <= N && N <= 5 {
		fmt.Println("Not Weird")
	} else if 6 <= N && N <= 20 {
		fmt.Println("Weird")
	} else if N > 20 {
		fmt.Println("Not Weird")
	}
}
