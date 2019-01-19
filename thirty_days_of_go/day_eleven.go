package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	var max int32 = math.MinInt32
	for x := 0; x <= 3; x++ {
		for y := 0; y <= 3; y++ {
			sum := calcHourglassSum(x, y, arr)
			if sum > max {
				max = sum
			}
		}
	}
	fmt.Println(max)
}

func calcHourglassSum(x int, y int, arr [][]int32) int32 {
	return arr[x][y] + arr[x][y+1] + arr[x][y+2] +
		arr[x+1][y+1] +
		arr[x+2][y] + arr[x+2][y+1] + arr[x+2][y+2]
}
