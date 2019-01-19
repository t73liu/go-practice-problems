package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nk := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nk[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		kTemp, err := strconv.ParseInt(nk[1], 10, 64)
		checkError(err)
		k := int32(kTemp)

		max := int32(0)
		for i := int32(1); i <= n; i++ {
			if max == k-1 {
				break
			}
			for j := i + 1; j <= n; j++ {
				curr := i & j
				if max < curr && curr < k {
					max = curr
				}
				if max == k-1 {
					break
				}
			}
		}
		fmt.Println(max)
	}
}
