package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	entries, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	phone := map[string]int64{}

	for i := int64(0); i < entries; i++ {
		scanner.Scan()
		arr := strings.Split(scanner.Text(), " ")
		phone[arr[0]], _ = strconv.ParseInt(arr[1], 10, 64)
	}

	for scanner.Scan() {
		name := scanner.Text()
		val, ok := phone[name]
		if ok {
			fmt.Printf("%s=%d\n", name, val)
		} else {
			fmt.Println("Not found")
		}
	}
}
