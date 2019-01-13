package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertToIntArray(arr []string) []int64 {
	vsm := make([]int64, len(arr))
	for i, v := range arr {
		vsm[i], _ = strconv.ParseInt(v, 10, 64)
	}
	return vsm
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	size, _ := strconv.ParseInt(readLine(reader), 10, 64)
	arr := convertToIntArray(strings.Split(readLine(reader), " "))
	swaps := 0

	for i := int64(0); i < size; i++ {
		currSwaps := 0

		for j := int64(0); j < size-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				currSwaps += 1
			}
		}

		swaps += currSwaps
		if currSwaps == 0 {
			break
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", swaps)
	fmt.Printf("First Element: %d\n", arr[0])
	fmt.Printf("Last Element: %d\n", arr[size-1])
}
