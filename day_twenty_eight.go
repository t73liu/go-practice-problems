package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)

	var arr []string
	for NItr := 0; NItr < int(N); NItr++ {
		firstNameEmailID := strings.Split(readLine(reader), " ")

		firstName := firstNameEmailID[0]

		emailID := firstNameEmailID[1]

		if strings.HasSuffix(emailID, "@gmail.com") {
			arr = append(arr, firstName)
		}
	}
	sort.Strings(arr)
	for _, name := range arr {
		fmt.Println(name)
	}
}
