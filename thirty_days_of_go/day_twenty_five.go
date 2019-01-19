package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	loops, _ := strconv.ParseInt(readLine(reader), 10, 64)

	for i := int64(0); i < loops; i++ {
		n, _ := strconv.ParseInt(readLine(reader), 10, 64)
		if n == 1 {
			fmt.Println("Not prime")
		} else if n <= 3 {
			fmt.Println("Prime")
		} else {
			prime := true
			for div := int64(2); div*div <= n; div++ {
				if n%div == 0 {
					prime = false
					break
				}
			}
			if prime {
				fmt.Println("Prime")
			} else {
				fmt.Println("Not prime")
			}
		}
	}
}
